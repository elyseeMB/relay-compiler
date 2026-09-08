package securecookie

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
)

type Config struct {
	Name string

	Secret string

	Domain string

	Path string

	MaxAge int

	Secure bool

	HTTPOnly bool

	SameSite http.SameSite
}

func DefaultConfig(name, secret string) Config {
	return Config{
		Name:   name,
		Secret: secret,

		Path:     "/",
		MaxAge:   86400 * 30,
		Secure:   true,
		HTTPOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}

func Set(w http.ResponseWriter, config Config, value string) error {
	signedValue, err := Sign(value, config.Secret)

	if err != nil {
		return fmt.Errorf("failed to sign cookie value %w", err)

	}

	cookie := &http.Cookie{
		Name:     config.Name,
		Value:    signedValue,
		Path:     config.Path,
		Domain:   config.Domain,
		MaxAge:   config.MaxAge,
		Secure:   config.Secure,
		HttpOnly: config.HTTPOnly,
		SameSite: config.SameSite,
	}

	http.SetCookie(w, cookie)
	return nil
}

func Sign(value, secret string) (string, error) {

	if secret == "" {
		return "", fmt.Errorf("secret cannot be empty")
	}

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(value))

	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	return value + "." + signature, nil

}
