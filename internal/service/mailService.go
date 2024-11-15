package service

import (
	"io"
	"mime/multipart"

	"github.com/exoneges/doodocs-days-backend/internal/config"
)

func AnalyzeMailFile(file multipart.File, filename string) ([]byte, error) {

	mimeType := detectMimeType(file, filename)

	switch mimeType {
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document", "application/pdf":
		// Process futher
	default:
		return nil, config.ErrFormatNotSupported
	}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
