package handler

import (
	"fmt"
	"net/http"

	"github.com/exoneges/doodocs-days-backend/internal/config"
	"github.com/exoneges/doodocs-days-backend/internal/utils"
)

func BasicAuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := config.UpdateENV()
		if err != nil {
			utils.SendJSONError(w, http.StatusInternalServerError, err, "server-side error, contact admin to enable authorization")
			return
		}
		// Extract username and password from the request's Basic Auth header
		username, password, ok := r.BasicAuth()
		// Check if Basic Auth is provided and if the credentials match the configured admin credentials
		if !ok || username != config.ENV_AUTH_USER || password != config.ENV_AUTH_PASS {
			utils.LogRequest(r, "Middleware check")
			// If authentication fails, set the WWW-Authenticate header
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			utils.SendJSONError(w, http.StatusUnauthorized, fmt.Errorf("authentication failed"), "Authentication failed")
			return
		}

		// If authentication succeeds, call the next handler
		next.ServeHTTP(w, r)
	})
}
