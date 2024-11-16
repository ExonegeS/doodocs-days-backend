package config

import "fmt"

var (
	ErrMimeNotSupported   error = fmt.Errorf("this mimetype is not supported. Only: application/vnd.openxmlformats-officedocument.wordprocessingml.document | application/xml | image/jpeg | image/png")
	ErrFormatNotSupported error = fmt.Errorf("provided file has unsupported format")
	ErrInvalidZipFile     error = fmt.Errorf("provided file is not ZIP archive")
	ErrWrongArraySize     error = fmt.Errorf("unexpected or wrong size of array")
	ErrEmptyFile          error = fmt.Errorf("file data is empty or file is missing")
	ErrCorruptedFileData  error = fmt.Errorf("file data is empty or corrupted")
)
