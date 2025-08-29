package api

import (
	"errors"
	"net/http"

	console_v1 "github.com/elyseeMB/relay-compiler/pkg/server/api/console/v1"
	"github.com/elyseeMB/relay-compiler/pkg/usrmgr"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.gearno.de/kit/httpserver"
	"go.gearno.de/kit/log"
)

type (
	Config struct {
		AllowedOrigins []string
		Usrmgr         *usrmgr.Service
		Logger         *log.Logger
	}

	Server struct {
		cfg Config
	}
)

var (
	ErrMissingTPServic      = errors.New("server configuration requires a valid tp.Service instance")
	ErrMissingUsrmgrService = errors.New("server configuration requires a valid usrmgr.Service instance")
)

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	httpserver.RenderJSON(
		w,
		http.StatusMethodNotAllowed,
		map[string]string{
			"error": "method not allowed",
		},
	)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	httpserver.RenderJSON(
		w,
		http.StatusNotFound,
		map[string]string{
			"error": "not found",
		},
	)
}

func NewServer(cfg Config) (*Server, error) {
	return &Server{
		cfg: cfg,
	}, nil

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	corsOpts := cors.Options{
		AllowedOrigins: s.cfg.AllowedOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowedHeaders: []string{
			"content-type", "traceparent", "authorization",
		},
		ExposedHeaders: []string{
			"x-Request-id",
		},
		AllowCredentials:   true,
		OptionsPassthrough: false,
		Debug:              false,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(corsOpts))

	cfgMux := console_v1.ConfigNewMux{
		Logger:    s.cfg.Logger.Named("console.v1"),
		UsrmgrSvc: s.cfg.Usrmgr,
	}

	router.Mount(
		"/console/v1",
		console_v1.NewMux(&cfgMux),
	)

	router.ServeHTTP(w, r)

}
