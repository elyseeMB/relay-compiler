package console_v1

import (
	"encoding/json"
	"log"
	"net/http"

	"go.gearno.de/kit/httpserver"
)

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		FullName string `json:"Fullname"`
		Role     string `json:"role"`
	}
)

func SignUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req SignUpRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return
		}

		log.Printf("Requête SignUp: %+v", req)

		httpserver.RenderJSON(w, http.StatusOK, map[string]string{
			"server": "voiture",
		})

	}

}
