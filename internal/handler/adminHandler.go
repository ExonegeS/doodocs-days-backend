package handler

import (
	"net/http"
	"os"

	"github.com/exoneges/doodocs-days-backend/internal/config"
	"github.com/exoneges/doodocs-days-backend/internal/utils"
)

func SetMuxAdminHanle(mux *http.ServeMux) {

	mux.Handle("/api/admin/log", BasicAuthMiddleware(http.HandlerFunc(sendLogs)))
	mux.Handle("/api/admin/log/", BasicAuthMiddleware(http.HandlerFunc(sendLogs)))
}

func sendLogs(w http.ResponseWriter, r *http.Request) {

	utils.LogRequest(r, "")
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	logData, err := os.ReadFile(config.DIR_LOGGER)
	if err != nil {
		utils.SendJSONError(w, http.StatusInternalServerError, err, "Cannot read data from app.log")
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(logData)
}
