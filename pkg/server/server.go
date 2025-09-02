package server

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/elyseeMB/relay-compiler/pkg/server/api"
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

	// Assets (JS/CSS) routes
	s.router.HandleFunc("/assets/*", viteAssets.ServeAssets)

	// Catch-all pour les routes frontend
	s.router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		// Pour la route racine, utilise le handler React
		if r.URL.Path == "/" {
			frontMiddleware(web.HomeHandler)(w, r)
			return
		}

		// Pour les autres fichiers statiques, essaie de les servir
		// Si le fichier n'existe pas, sert l'app React (SPA routing)
		if strings.HasPrefix(r.URL.Path, "/assets/") {
			// Les assets sont déjà gérés plus haut
			http.NotFound(w, r)
			return
		}

		// Essaie de servir le fichier statique
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
