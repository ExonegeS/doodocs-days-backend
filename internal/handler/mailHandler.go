package handler

import (
	"net/http"

	"github.com/exoneges/doodocs-days-backend/internal/utils"
)

func SetMuxMailHanle(mux *http.ServeMux) {
	mux.Handle("/api/mail/file", BasicAuthMiddleware(http.HandlerFunc(postMailFile)))
	mux.Handle("/api/mail/file/", BasicAuthMiddleware(http.HandlerFunc(postMailFile)))
}

func postMailFile(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r, "")
}
