// Code AutoGenerated; DO NOT EDIT.

package types

// Sticker Represents a sticker.
type Sticker struct {
	Emoji            string        `json:"emoji,omitempty"`
	FileID           string        `json:"file_id"`
	FileSize         int           `json:"file_size,omitempty"`
	FileUniqueID     string        `json:"file_unique_id"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	SetName          string        `json:"set_name,omitempty"`
	Thumb            *PhotoSize    `json:"thumb,omitempty"`
	Width            int64         `json:"width"`
}
