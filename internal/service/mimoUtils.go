package service

import (
	"archive/zip"
	"io"
	"net/http"
)

func detectZipMimeType(file *zip.File) string {

	// Open the file to read its content
	f, err := file.Open()
	if err != nil {
		return "unknown"
	}
	defer f.Close()

	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil && err != io.EOF {
		return "unknown"
	}

	return http.DetectContentType(buffer)
}
