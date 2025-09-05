package db

type (
	authConfig struct {
		Cookie        cookieConfig   `json:"cookie"`
		Password      passwordConfig `json:"password"`
		DisableSignup bool           `json:"disable-signup"`
	}

	cookieConfig struct {
		Domain   string `json:"domain"`
		Secret   string `json:"secret"`
		Duration int    `json:"duration"`
		Name     string `json:"name"`
	}

	passwordConfig struct {
		Iterations uint32 `json:"iterations"`
		Pepper     string `json:"pepper"`
	}
)
