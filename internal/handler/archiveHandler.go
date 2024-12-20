package handler

import (
	"encoding/json"
	"net/http"

	"github.com/exoneges/doodocs-days-backend/internal/config"
	"github.com/exoneges/doodocs-days-backend/internal/service"
	"github.com/exoneges/doodocs-days-backend/internal/utils"
	"github.com/exoneges/doodocs-days-backend/models"
)

func SetMuxArchiveHanle(mux *http.ServeMux) {

	mux.Handle("/api/archive/information", http.HandlerFunc(postArchiveInformation))
	mux.Handle("/api/archive/information/", http.HandlerFunc(postArchiveInformation))

	mux.Handle("/api/archive/files", http.HandlerFunc(postArchiveFiles))
	mux.Handle("/api/archive/files/", http.HandlerFunc(postArchiveFiles))
}

func postArchiveInformation(w http.ResponseWriter, r *http.Request) {

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
	defer file.Close()

	// Delegate processing to the service layer
	archiveData, err := service.AnalyzeZipFile(models.FileWithMeta{
		File:        file,
		Filename:    header.Filename,
		ContentType: header.Header.Get("Content-Type")})

	if err != nil {
		switch err {
		case config.ErrInvalidZipFile:
			utils.SendJSONError(w, http.StatusBadRequest, err, err.Error())
		case config.ErrMimeNotSupported:
			utils.SendJSONError(w, http.StatusNotAcceptable, err, err.Error())
		default:
			utils.SendJSONError(w, http.StatusInternalServerError, err, "Failed to analyze ZIP file")
		}
		return
	}

	// Return the analysis result as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(archiveData)
}

func postArchiveFiles(w http.ResponseWriter, r *http.Request) {

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

	files := make([]models.FileWithMeta, 0)

	formFiles, ok := r.MultipartForm.File["files[]"]
	if !ok {
		utils.SendJSONError(w, http.StatusBadRequest, nil, "No files found in the request")
		return
	}
	for _, fileHeader := range formFiles {
		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			utils.SendJSONError(w, http.StatusInternalServerError, err, "Failed to open file")
			return
		}
		defer file.Close()
		files = append(files, models.FileWithMeta{
			File:        file,
			Filename:    fileHeader.Filename,
			ContentType: fileHeader.Header.Get("Content-Type")})
	}

	// Delegate processing to the service layer
	archiveData, err := service.ConstructArchive(files)
	if err != nil {
		switch err {
		case config.ErrCorruptedFileData:
			utils.SendJSONError(w, http.StatusBadRequest, err, err.Error())
		case config.ErrMimeNotSupported:
			utils.SendJSONError(w, http.StatusBadRequest, err, err.Error())
		default:
			utils.SendJSONError(w, http.StatusInternalServerError, err, "failed to analyze ZIP file")
		}
		return
	}

	// Return the result ZIP file as binary data
	w.Header().Set("Content-Type", "application/zip")
	w.Write(archiveData)
}
