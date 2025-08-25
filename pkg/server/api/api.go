package api

import (
	"errors"
	"net/http"

	console_v1 "github.com/elyseeMB/relay-compiler/pkg/server/api/console/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.gearno.de/kit/httpserver"
)

type (
	Config struct {
		AllowedOrigins []string
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

	router.Mount(
		"/console/v1",
		console_v1.NewMux(),
	)

	router.ServeHTTP(w, r)

}
