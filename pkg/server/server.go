package server

import (
	"net/http"
	"strings"

	"github.com/elyseeMB/relay-compiler/pkg/server/api"
	"github.com/elyseeMB/relay-compiler/pkg/usrmgr"
	"github.com/go-chi/chi/v5"
	"go.gearno.de/kit/log"
)

type Server struct {
	apiServer *api.Server
	router    *chi.Mux
}

type Config struct {
	AllowedOrigins []string
	Usrmgr         *usrmgr.Service
	Logger         *log.Logger
}

func NewServer(cfg Config) (*Server, error) {
	apiCfg := api.Config{
		AllowedOrigins: cfg.AllowedOrigins,
		Usrmgr:         cfg.Usrmgr,
		Logger:         cfg.Logger.Named("api"),
	}

	apiServer, err := api.NewServer(apiCfg)

	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()

	server := &Server{
		apiServer: apiServer,
		router:    router,
	}

	// Setup Routes
	server.setupRoutes()

	return server, nil
}

func (s *Server) setupRoutes() {

	s.router.Mount("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
		if r.URL.Path == "" {
			r.URL.Path = "/"
		}
		s.apiServer.ServeHTTP(w, r)
	}))

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
