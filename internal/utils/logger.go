package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/exoneges/doodocs-days-backend/internal/config"
)

func LogRequest(r *http.Request, addInfo string) {
	if config.PROTECTED {
		config.LOGGER.Info("Request",
			"From", "x.x.x.x",
			"Method", r.Method,
			"URL", r.URL,
			"Additional", addInfo)
		slog.Info("Request",
			"From", "x.x.x.x",
			"Method", r.Method,
			"URL", r.URL,
			"Additional", addInfo)
	} else {
		config.LOGGER.Info("Request",
			"From", r.RemoteAddr,
			"Method", r.Method,
			"URL", r.URL,
			"Additional", addInfo)
		slog.Info("Request",
			"From", r.RemoteAddr,
			"Method", r.Method,
			"URL", r.URL,
			"Additional", addInfo)
	}
}

func SendJSONError(w http.ResponseWriter, code int, err error, message string) {
	config.LOGGER.Error(err.Error())
	slog.Error(err.Error())
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
