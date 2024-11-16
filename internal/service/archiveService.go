package service

import (
	"archive/zip"
	"bytes"
	"io"

	"github.com/exoneges/doodocs-days-backend/internal/config"
	"github.com/exoneges/doodocs-days-backend/models"
)

func AnalyzeZipFile(zipFIle models.FileWithMeta) (models.Archive, error) {

	// Check filename suffix
	if zipFIle.ContentType != "application/zip" {
		return models.Archive{}, config.ErrInvalidZipFile
	}

	fileData, err := io.ReadAll(zipFIle.File)
	if err != nil {
		return models.Archive{}, err
	}

	// Create a bytes.Reader for zip.NewReader
	reader := bytes.NewReader(fileData)
	zipReader, err := zip.NewReader(reader, int64(len(fileData)))
	if err != nil {
		return models.Archive{}, err
	}

	var totalSize float64
	var files []models.FileObject

	// Iterate through files in the ZIP
	for _, file := range zipReader.File {
		fileSize := float64(file.UncompressedSize64)
		mimeType := detectZipMimeType(file)
		switch mimeType {
		case "application/octet-stream": // Directory or empty file
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
		Filename:    zipFIle.Filename,
		ArchiveSize: float64(len(fileData)), // Actual ZIP archive size
		TotalSize:   totalSize,              // Total size of uncompressed files
		TotalFiles:  float64(len(files)),
		Files:       files,
	}, nil
}

func ConstructArchive(files []models.FileWithMeta) ([]byte, error) {
	// Create a buffer to write the archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)
	for _, file := range files {

		// Check that multipart.File is not empty file
		buffer := make([]byte, 1) // Check only the first byte
		n, _ := file.File.Read(buffer)
		file.File.Seek(0, io.SeekStart) // Reset the file pointer
		if n == 0 {
			return nil, config.ErrEmptyFile
		}
		switch file.ContentType {
		case "application/octet-stream":
			w.Close()
			return nil, config.ErrCorruptedFileData
		case "image/png", "image/jpeg", "application/xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.document":

			// Add the file to the archive.
		default:
			w.Close()
			return nil, config.ErrMimeNotSupported
		}

		// Create a new file in the ZIP archive.
		f, err := w.Create(file.Filename)
		if err != nil {
			return nil, err
		}

		// Read the content of the `multipart.File`.
		fileContent, err := io.ReadAll(file.File)
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
