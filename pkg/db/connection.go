package db

import (
	"context"
	"fmt"

	"github.com/elyseeMB/relay-compiler/pkg/coredata"
	"github.com/prometheus/client_golang/prometheus"
	"go.gearno.de/kit/log"
	"go.gearno.de/kit/migrator"
	"go.gearno.de/kit/pg"
	"go.gearno.de/kit/unit"
	"go.opentelemetry.io/otel/trace"
)

type (
	Implm struct {
		cfg config
	}

	config struct {
		Hostname string   `json:hostname`
		Pg       pgConfig `json:pg`
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
		},
	}
}

func (impl *Implm) GetConfiguration() any {
	return &impl.cfg
}

func (impl *Implm) Run(parentCtx context.Context, l *log.Logger, r prometheus.Registerer, tp trace.TracerProvider) error {

	pgClient, err := pg.NewClient(
		impl.cfg.Pg.Options(pg.WithLogger(l), pg.WithRegisterer(r), pg.WithTracerProvider(tp))...,
	)
	if err != nil {
		return fmt.Errorf("Cannot create pg client: %w", err)
	}

	// Exécuter les migrations
	err = migrator.NewMigrator(pgClient, coredata.Migrations, l.Named("migrations")).Run(parentCtx, "migrations")
	if err != nil {
		return fmt.Errorf("cannot migrate database schema: %w", err)
	}

	defer pgClient.Close()
	return context.Cause(parentCtx)

}
