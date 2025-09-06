package console_v1

import (
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/elyseeMB/relay-compiler/pkg/server/api/console/v1/schema"
	"github.com/elyseeMB/relay-compiler/pkg/server/api/console/v1/types"
	"github.com/elyseeMB/relay-compiler/pkg/usrmgr"
	"github.com/go-chi/chi/v5"
	"github.com/vektah/gqlparser/v2/ast"
	"go.gearno.de/kit/log"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type (
	AuthConfig struct {
		CookieName      string
		CookieDomain    string
		SessionDuration time.Duration
		CokkieSecret    string
	}

	Resolver struct {
		Articles []*types.Article
		Users    []*types.User
	}

	ConfigNewMux struct {
		AuthConfig AuthConfig
		Logger     *log.Logger
		UsrmgrSvc  *usrmgr.Service
	}

	SignUpConfig struct {
		AuthConfig AuthConfig
		UsrmgrSvc  *usrmgr.Service
	}
)

const defaultPort = "8080"

func NewMux(cfg *ConfigNewMux) *chi.Mux {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewMux()

	r.Post("/auth/register", SignUpHandler(SignUpConfig{
		AuthConfig: cfg.AuthConfig,
		UsrmgrSvc:  cfg.UsrmgrSvc,
	}))

	r.Get("/", playground.Handler("GraphQL", "/api/console/v1/query"))
	r.Post("/query", graphqlHandler())

	return r

}

func graphqlHandler() http.HandlerFunc {

	srv := handler.New(schema.NewExecutableSchema(schema.Config{Resolvers: &Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv.ServeHTTP

}
