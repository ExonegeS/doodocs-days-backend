package service

import (
	"archive/zip"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
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

func detectMimeType(file multipart.File, fileName string) string {

	// Check file extension as a fallback if MIME type is not detected properly
	if strings.HasSuffix(fileName, ".json") {
		return "application/json"
	}

	// Seek to the beginning of the file
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return "unknown"
	}

	// Read the first 512 bytes for MIME type detection
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return "unknown"
	}

	// Reset the file pointer to the start for further operations
	_, _ = file.Seek(0, io.SeekStart)

	// Detect the MIME type
	mimeType := http.DetectContentType(buffer)

	// If it returns "application/octet-stream", check if it's a JSON file by extension
	if mimeType == "application/octet-stream" {
		// If mimeType detection is not Success try to check the filename extention
		switch {
		case strings.HasSuffix(fileName, ".png"):
			return "image/png"
		case strings.HasSuffix(fileName, ".jpeg") || strings.HasSuffix(fileName, ".jpg"):
			return "image/jpeg"
		case strings.HasSuffix(fileName, ".xml"):
			return "application/xml"
		case strings.HasSuffix(fileName, ".docx") || strings.HasSuffix(fileName, ".doc"):
			return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
		}
	}
	return mimeType
}

func validateFileContent(file multipart.File, mimeType string) error {

	_ = file
	_ = mimeType
	// TODO : check multipart.File contents to match the mimeType and be readable
	// return config.ErrCorruptedFileData

	return nil
}
