package models

type Archive struct {
	Filename     string     `json:"filename" xml:"Filename"`
	Archive_size float64    `json:"archive_size" xml:"Archive_size"`
	Total_size   float64    `json:"total_size" xml:"Total_size"`
	Total_files  float64    `json:"total_files" xml:"Total_files"`
	Files        FileObject `json:"files" xml:"Files"`
}

type FileObject struct {
	File_path string  `json:"file_path" xml:"Path"`
	Size      float64 `json:"size" xml:"size"`
	Mimetype  string  `json:"mimetype" xml:"Mimetype"`
}
