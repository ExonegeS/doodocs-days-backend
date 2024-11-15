package handler

import (
	"net/http"
	"os"
)

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract username and password from the request's Basic Auth header
		username, password, ok := r.BasicAuth()
		// Check if Basic Auth is provided and if the credentials match the configured admin credentials
		if !ok || username != os.Getenv("DOODOCS_DAYS2_BACKEND_AUTH_USERNAME") || password != os.Getenv("DOODOCS_DAYS2_BACKEND_AUTH_PASSWORD") {
			// If authentication fails, set the WWW-Authenticate header
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			// Return a 401 Unauthorized status
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authentication succeeds, call the next handler
		next.ServeHTTP(w, r)
	})
}
