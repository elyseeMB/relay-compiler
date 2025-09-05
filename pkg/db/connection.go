package db

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/elyseeMB/relay-compiler/pkg/crypto/passwdhash"
	"github.com/elyseeMB/relay-compiler/pkg/server"
	console_v1 "github.com/elyseeMB/relay-compiler/pkg/server/api/console/v1"
	"github.com/elyseeMB/relay-compiler/pkg/usrmgr"
	"github.com/prometheus/client_golang/prometheus"
	"go.gearno.de/kit/httpserver"
	"go.gearno.de/kit/log"
	"go.gearno.de/kit/pg"
	"go.gearno.de/kit/unit"
	"go.opentelemetry.io/otel/trace"
)

type (
	Implm struct {
		cfg config
	}

	config struct {
		Hostname string     `json:"hostname"`
		Pg       pgConfig   `json:"pg"`
		Auth     authConfig `json:"auth"`
	}
)

var (
	_ unit.Configurable = (*Implm)(nil)
	_ unit.Runnable     = (*Implm)(nil)
)

func New() *Implm {
	return &Implm{
		cfg: config{
			Hostname: "localhost:8080",
			Pg: pgConfig{
				Username: "tp",
				Password: "password",
				Database: "tp_database",
			},
			Auth: authConfig{
				Password: passwordConfig{
					Pepper:     "this-is-a-secure-pepper",
					Iterations: 1000000,
				},
				Cookie: cookieConfig{
					Name:     "SSID",
					Secret:   "this-is-a-secure-cookie",
					Duration: 24,
					Domain:   "localhost",
				},
				DisableSignup: false,
			},
		},
	}
}

const defaultPort = "8080"

func (impl *Implm) GetConfiguration() any {
	return &impl.cfg
}

func (impl *Implm) Run(parentCtx context.Context, l *log.Logger, r prometheus.Registerer, tp trace.TracerProvider) error {

	tracer := tp.Tracer("react-relay")

	ctx, rootSpan := tracer.Start(parentCtx, "connection.Run")
	defer rootSpan.End()

	pgClient, err := pg.NewClient(
		impl.cfg.Pg.Options(pg.WithLogger(l), pg.WithRegisterer(r), pg.WithTracerProvider(tp))...,
	)
	if err != nil {
		return fmt.Errorf("cannot create pg client: %w", err)
	}

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(context.Canceled)

	pepper := []byte("12345678901234567890123456789012")

	hp, err := passwdhash.NewProfile(pepper, uint32(800000))
	if err != nil {
		return fmt.Errorf("cannot create hashing profile: %w", err)
	}

	usrmgrService, err := usrmgr.NewService(ctx, pgClient, hp)

	// Exécuter les migrations
	// err = migrator.NewMigrator(pgClient, coredata.Migrations, l.Named("migrations")).Run(parentCtx, "migrations")
	// if err != nil {
	// 	return fmt.Errorf("cannot migrate database schema: %w", err)
	// }

	serverHandler, err := server.NewServer(server.Config{
		AllowedOrigins: []string{"http://localhost:5173"},
		Logger:         l.Named("http.Server"),
		Usrmgr:         usrmgrService,
		Auth: console_v1.AuthConfig{
			CookieName:      impl.cfg.Auth.Cookie.Name,
			CookieDomain:    impl.cfg.Auth.Cookie.Domain,
			SessionDuration: time.Duration(impl.cfg.Auth.Cookie.Duration) * time.Hour,
			CokkieSecret:    impl.cfg.Auth.Cookie.Secret,
		},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err != nil {
		return fmt.Errorf("cannot create server: %w", err)
	}

	apiServerCtx, stopApiServer := context.WithCancel(context.Background())

	defer stopApiServer()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := impl.runApiServer(apiServerCtx, l, r, tp, serverHandler); err != nil {
			cancel(fmt.Errorf("api server crashed: %w", err))
		}
	}()

	<-ctx.Done()

	stopApiServer()
	pgClient.Close()

	wg.Wait()
	return context.Cause(ctx)

}

func (impl *Implm) runApiServer(
	ctx context.Context,
	l *log.Logger,
	r prometheus.Registerer,
	tp trace.TracerProvider,
	handler http.Handler,
) error {
	tracer := tp.Tracer("github.com/getprobo/probo/pkg/probod")
	ctx, span := tracer.Start(ctx, "probod.runApiServer")
	defer span.End()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	addr := "localhost:" + port

	apiServer := httpserver.NewServer(
		addr,
		handler,
		httpserver.WithLogger(l),
		httpserver.WithRegisterer(r),
		httpserver.WithTracerProvider(tp),
	)

	l.Info("starting api server", log.String("addr", apiServer.Addr))
	span.AddEvent("API server starting")

	listener, err := net.Listen("tcp", apiServer.Addr)
	if err != nil {
		span.RecordError(err)
		return fmt.Errorf("cannot listen on %q: %w", apiServer.Addr, err)
	}
	defer listener.Close()

	serverErrCh := make(chan error, 1)
	go func() {
		err := apiServer.Serve(listener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrCh <- fmt.Errorf("cannot server http request: %w", err)
		}
		close(serverErrCh)
	}()

	l.Info("api server started")
	span.AddEvent("API server started")

	select {
	case err := <-serverErrCh:
		if err != nil {
			span.RecordError(err)
		}
		return err
	case <-ctx.Done():
	}

	l.InfoCtx(ctx, "shutting down api server")
	span.AddEvent("API server shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := apiServer.Shutdown(shutdownCtx); err != nil {
		span.RecordError(err)
		return fmt.Errorf("cannot shutdown api server: %w", err)
	}

	span.AddEvent("API server shutdown complete")
	return ctx.Err()
}
