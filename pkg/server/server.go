package server

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/elyseeMB/relay-compiler/pkg/server/api"
	console_v1 "github.com/elyseeMB/relay-compiler/pkg/server/api/console/v1"
	"github.com/elyseeMB/relay-compiler/pkg/server/web"
	"github.com/elyseeMB/relay-compiler/pkg/usrmgr"
	"github.com/go-chi/chi/v5"
	"go.gearno.de/kit/log"
)

//go:embed all:public
var assets embed.FS

type Server struct {
	apiServer *api.Server
	router    *chi.Mux
}

type Config struct {
	AllowedOrigins []string
	Usrmgr         *usrmgr.Service
	Logger         *log.Logger
	Auth           console_v1.AuthConfig
}

func NewServer(cfg Config) (*Server, error) {
	apiCfg := api.Config{
		AllowedOrigins: cfg.AllowedOrigins,
		Usrmgr:         cfg.Usrmgr,
		Auth:           cfg.Auth,
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
	publicFS, err := fs.Sub(assets, "public")
	if err != nil {
		panic(fmt.Sprintf("Cannot sub public directory %v", err))
	}

	viteAssets := web.NewViteAssets(publicFS)
	frontMiddleware := createFrontEndMiddleware(*viteAssets)
	publicServer := http.FileServer(http.FS(publicFS))

	// API routes
	s.router.Mount("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
		if r.URL.Path == "" {
			r.URL.Path = "/"
		}
		s.apiServer.ServeHTTP(w, r)
	}))

	s.router.HandleFunc("/assets/*", viteAssets.ServeAssets)

	s.router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			frontMiddleware(web.HomeHandler)(w, r)
			return
		}

		if strings.HasPrefix(r.URL.Path, "/assets/") {
			http.NotFound(w, r)
			return
		}

		publicServer.ServeHTTP(w, r)
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Inject assets tags (as a string) in the context
func createFrontEndMiddleware(vite web.ViteAssets) func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	html := vite.GetHeadHTML()
	return func(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "assets", html)
			next(w, r.WithContext(ctx))
		}
	}
}
