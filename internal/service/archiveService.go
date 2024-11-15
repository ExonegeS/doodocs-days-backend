package service

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/exoneges/doodocs-days-backend/internal/config"
	"github.com/exoneges/doodocs-days-backend/models"
)

func AnalyzeZipFile(zipFIle multipart.File, filename string) (models.Archive, error) {

	// Check filename suffix
	if !strings.HasSuffix(filename, ".zip") {
		return models.Archive{}, config.ErrInvalidZipFile
	}

	fileData, err := io.ReadAll(zipFIle)
	if err != nil {
		return models.Archive{}, fmt.Errorf("failed to read uploaded file: %w", err)
	}

	// Create a bytes.Reader for zip.NewReader
	reader := bytes.NewReader(fileData)
	zipReader, err := zip.NewReader(reader, int64(len(fileData)))
	if err != nil {
		return models.Archive{}, fmt.Errorf("unable to read zip file: %w", err)
	}

	var totalSize float64
	var files []models.FileObject

	// Iterate through files in the ZIP
	for _, file := range zipReader.File {
		fileSize := float64(file.UncompressedSize64)
		mimeType := detectZipMimeType(file)
		switch mimeType {
		case "application/octet-stream": // Directory
		// case "image/png", "image/jpeg", "application/xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		// 	files = append(files, models.FileObject{
		// 		FilePath: file.Name,
		// 		Size:     fileSize,
		// 		MimeType: mimeType,
		// 	})
		// default:
		// 	return models.Archive{}, config.ErrMimeNotSupported
		default:
			files = append(files, models.FileObject{
				FilePath: file.Name,
				Size:     fileSize,
				MimeType: mimeType,
			})
		}

		totalSize += fileSize
	}
	return models.Archive{
		Filename:    filename,
		ArchiveSize: float64(len(fileData)), // Actual ZIP archive size
		TotalSize:   totalSize,              // Total size of uncompressed files
		TotalFiles:  float64(len(files)),
		Files:       files,
	}, nil
}

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

func ConstructArchive(files []multipart.File, fileNames []string) ([]byte, error) {
	if len(files) != len(fileNames) {
		return nil, config.ErrWrongArraySize
	}

	// Create a buffer to write the archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)
	for i, file := range files {

		// Check that multipart.File is not empty file
		buffer := make([]byte, 1) // Check only the first byte
		n, _ := file.Read(buffer)
		file.Seek(0, io.SeekStart) // Reset the file pointer
		if n == 0 {
			return nil, config.ErrEmptyFile
		}

		// Could use here basic if statement, but nah, boring
		mimeType := detectMimeType(file)
		fmt.Println(mimeType, fileNames[i])
		switch mimeType {
		case "application/octet-stream":
			w.Close()
			return nil, config.ErrCorruptedFileData
		case "image/png", "image/jpeg", "application/xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
			err := validateFileContent(file, mimeType)
			if err != nil {
				return nil, err
			}
			// Add the file to the archive.
		default:
			w.Close()
			return nil, config.ErrMimeNotSupported
		}

		// Create a new file in the ZIP archive.
		f, err := w.Create(fileNames[i])
		if err != nil {
			return nil, err
		}

		// Read the content of the `multipart.File`.
		fileContent, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}

		// Write the content to the archive.
		_, err = f.Write(fileContent)
		if err != nil {
			return nil, err
		}
	}

	// Close the ZIP writer to finalize the archive.
	err := w.Close()
	if err != nil {
		return nil, err
	}

	// Return the constructed archive as a byte slice.
	return buf.Bytes(), nil
}

func detectMimeType(file multipart.File) string {
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
	return http.DetectContentType(buffer)
}

func validateFileContent(file multipart.File, mimeType string) error {
	_ = file
	_ = mimeType
	// TODO : check multipart.File contents to match the mimeType and be readable
	// return config.ErrCorruptedFileData

	return nil
}
