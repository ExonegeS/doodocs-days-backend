package models

type Archive struct {
	Filename    string       `json:"filename" xml:"Filename"`
	ArchiveSize float64      `json:"archive_size" xml:"Archive_size"`
	TotalSize   float64      `json:"total_size" xml:"Total_size"`
	TotalFiles  float64      `json:"total_files" xml:"Total_files"`
	Files       []FileObject `json:"files" xml:"Files"`
}

type FileObject struct {
	FilePath string  `json:"file_path" xml:"Path"`
	Size     float64 `json:"size" xml:"size"`
	MimeType string  `json:"mimetype" xml:"Mimetype"`
}
