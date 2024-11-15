package service

import (
	"io"
	"mime/multipart"

	"github.com/exoneges/doodocs-days-backend/internal/config"
)

func AnalyzeMailFile(file multipart.File, filename, contentType string) ([]byte, error) {
	switch contentType {
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

func AnalyzeMailReceivers(file multipart.File, filename, contentType string) ([]byte, error) {
	switch contentType {
	case "text/plain":
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
