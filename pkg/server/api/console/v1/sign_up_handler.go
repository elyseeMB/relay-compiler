package console_v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/elyseeMB/relay-compiler/pkg/usrmgr"
	"go.gearno.de/kit/httpserver"
)

type (
	SignUpRequest struct {
		Password string `json:"password"`
		FullName string `json:"Fullname"`
	}
)

func SignUpHandler(usermgrSvc *usrmgr.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req SignUpRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return
		}

		user, err := usermgrSvc.SignUp(r.Context(), req.FullName, req.Password)

		if err != nil {
			httpserver.RenderError(w, http.StatusBadRequest, fmt.Errorf("connot register user: %w", err))

			panic(fmt.Errorf("cannot register %w", err))

		}

		log.Printf("Requête SignUp: %+v", req)

		httpserver.RenderJSON(w, http.StatusOK, map[string]interface{}{
			"data": user,
		})

	}

}
