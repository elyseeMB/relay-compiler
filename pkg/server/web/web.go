package web

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/elyseeMB/relay-compiler/templates"
)

type viteManifestItem struct {
	File    string   `json:"file"`
	Name    string   `json:"name"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	CSS     []string `json:"css"`
}

type viteManifestData map[string]viteManifestItem

type ViteAssets struct {
	publicPath   string
	assets       fs.FS
	hasManifest  bool
	port         int16
	manifestData viteManifestData
}

func NewViteAssets(filesystem fs.FS) *ViteAssets {
	var data viteManifestData
	manifestPath := "assets/.vite/manifest.json"
	hasManifest := false

	f, err := filesystem.Open(manifestPath)
	if err == nil {
		defer f.Close()
		hasManifest = true

		// Parse le contenu du manifest
		decoder := json.NewDecoder(f)
		if err := decoder.Decode(&data); err != nil {
			fmt.Printf("Erreur lors du parsing du manifest: %v\n", err)
			hasManifest = false
			data = make(viteManifestData)
		}
	} else {
		fmt.Printf("Manifest non trouvé (mode dev): %v\n", err)
		data = make(viteManifestData)
	}

	return &ViteAssets{
		publicPath:   "/assets/",
		assets:       filesystem,
		hasManifest:  hasManifest,
		port:         5173,
		manifestData: data,
	}
}

func (v ViteAssets) ServeAssets(w http.ResponseWriter, r *http.Request) {
	if v.hasManifest {
		http.FileServer(http.FS(v.assets)).ServeHTTP(w, r)
		return
	}

	// Proxy everything to vite in dev mode
	u := *r.URL
	u.Host = fmt.Sprintf("%s:%d", strings.Split(r.Host, ":")[0], v.port)
	u.Scheme = "http"
	w.Header().Set("Location", u.String())
	w.WriteHeader(301)
}

func (v ViteAssets) GetHeadHTML() string {
	var sb strings.Builder

	if !v.hasManifest {
		// Mode développement
		sb.WriteString(fmt.Sprintf(`<script type="module" src="http://localhost:%[1]d/@vite/client"></script>
			<script src="http://localhost:%[1]d/assets/main.tsx" type="module"></script>`, v.port))
		return sb.String()
	}

	// Mode production : utilise le manifest
	for _, item := range v.manifestData {
		if item.IsEntry {
			// Charge les CSS d'abord
			for _, css := range item.CSS {
				sb.WriteString(fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s%s\">\n", v.publicPath, css))
			}
			// Puis le JS
			sb.WriteString(fmt.Sprintf("<script type=\"module\" src=\"%s%s\"></script>\n", v.publicPath, item.File))
		}
	}

	return sb.String()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	component := templates.Layout("John")
	component.Render(r.Context(), w)
}
