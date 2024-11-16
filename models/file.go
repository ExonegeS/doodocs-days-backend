package models

import "mime/multipart"

type FileWithMeta struct {
	File        multipart.File `json:"file_data" xml:"Data"`
	Filename    string         `json:"filename" xml:"Filename"`
	ContentType string         `json:"content_type" xml:"ContentType"`
}
