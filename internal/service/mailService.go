package service

import (
	"io"

	"github.com/exoneges/doodocs-days-backend/internal/config"
	"github.com/exoneges/doodocs-days-backend/models"
)

func AnalyzeMailFile(file models.FileWithMeta) ([]byte, error) {
	switch file.ContentType {
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document", "application/pdf":
		// Process futher
	default:
		return nil, config.ErrFormatNotSupported
	}

	fileContent, err := io.ReadAll(file.File)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

func AnalyzeMailReceivers(file models.FileWithMeta) ([]byte, error) {
	switch file.ContentType {
	case "text/plain":
		// Process futher
	default:
		return nil, config.ErrFormatNotSupported
	}

	fileContent, err := io.ReadAll(file.File)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
