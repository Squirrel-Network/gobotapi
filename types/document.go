package types

type Document struct {
	FileId string `json:"file_id"`
	FileName string `json:"file_name,omitempty"`
	FileSize int `json:"file_size,omitempty"`
	FileUniqueId string `json:"file_unique_id"`
	MimeType string `json:"mime_type,omitempty"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
}