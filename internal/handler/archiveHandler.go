package handler

import (
	"fmt"
	"net/http"
)

func SetMuxArchiveHanle(mux *http.ServeMux) {
	mux.Handle("/api/archive/information", http.HandlerFunc(postArchiveInformation))
	mux.Handle("/api/archive/information/", http.HandlerFunc(postArchiveInformation))
}

func postArchiveInformation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	data := make([]byte, 0)
	_, err := r.Body.Read(data)
	if err != nil {
		http.Error(w, "Error body reading", http.StatusBadRequest)
		return
	}

	fmt.Println(data)
	fmt.Println("ALERT! POST request received.")
}
