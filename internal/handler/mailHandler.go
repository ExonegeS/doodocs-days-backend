package handler

import (
	"net/http"

	"github.com/exoneges/doodocs-days-backend/internal/service"
	"github.com/exoneges/doodocs-days-backend/internal/utils"
)

func SetMuxMailHanle(mux *http.ServeMux) {

	mux.Handle("/api/mail/file", BasicAuthMiddleware(http.HandlerFunc(postMailFile)))
	mux.Handle("/api/mail/file/", BasicAuthMiddleware(http.HandlerFunc(postMailFile)))
}

func postMailFile(w http.ResponseWriter, r *http.Request) {

	utils.LogRequest(r, "")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	if err := r.ParseMultipartForm(32 << 20); err != nil { // 32 MB limit
		utils.SendJSONError(w, http.StatusBadRequest, err, "Unable to parse form")
		return
	}

	// Retrieve the file from the form
	file, header, err := r.FormFile("file")
	if err != nil {
		utils.SendJSONError(w, http.StatusBadRequest, err, "Failed retrieving ZIP file")
		return
	}

	_, err = service.AnalyzeMailFile(file, header.Filename)
}
