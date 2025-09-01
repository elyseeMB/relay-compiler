package console_v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/elyseeMB/relay-compiler/pkg/coredata"
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
			httpserver.RenderError(w, http.StatusBadRequest, fmt.Errorf("cannot decode body: %w", err))
			return
		}

		user, err := usermgrSvc.SignUp(r.Context(), req.FullName, req.Password)

		if err != nil {
			var errUserAlreadyExists *coredata.ErrUserAlreadyExists

			if errors.As(err, &errUserAlreadyExists) {
				httpserver.RenderError(w, http.StatusBadRequest, fmt.Errorf("connot register user: %w", err))
				panic(fmt.Errorf("cannot register %w", err))
			}

			var errSignupDisabled *usrmgr.ErrSignupDisabled
			if errors.As(err, &errSignupDisabled) {
				httpserver.RenderError(w, http.StatusBadRequest, fmt.Errorf("cannot register user: %w", err))
				return
			}

			panic(fmt.Errorf("cannot register user: %w", err))

		}

		log.Printf("Requête SignUp: %+v", req)

		httpserver.RenderJSON(w, http.StatusOK, map[string]interface{}{
			"data": user,
		})

	}

}
