package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/exoneges/doodocs-days-backend/internal/config"
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
		utils.SendJSONError(w, http.StatusBadRequest, err, "Failed retrieving file to send")
		return
	}

	fileData, err := service.AnalyzeMailFile(file, header.Filename, header.Header.Get("Content-Type"))
	if err != nil {
		switch err {
		case config.ErrFormatNotSupported:
			utils.SendJSONError(w, http.StatusBadRequest, err, "Unsupported file format")
		default:
			utils.SendJSONError(w, http.StatusInternalServerError, err, "Failed to analyze file to send")
		}
		return
	}
	_ = fileData

	// Retrieve the emails from the form
	receiversFile, header, err := r.FormFile("emails")
	if err != nil {
		utils.SendJSONError(w, http.StatusBadRequest, err, "Failed retrieving receivers emails file")
		return
	}
	receiversData, err := io.ReadAll(receiversFile)
	if err != nil {
		utils.SendJSONError(w, http.StatusBadRequest, err, "Failed reading receivers emails file")
		return
	}

	fmt.Println(string(receiversData))
}
