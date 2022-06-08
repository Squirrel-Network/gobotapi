// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// UploadStickerFile Use this method to upload a .PNG file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times)
// Returns the uploaded File on success.
type UploadStickerFile struct {
	PngSticker rawTypes.InputFile        `json:"png_sticker,omitempty"`
	UserID     int64                     `json:"user_id"`
	Progress   rawTypes.ProgressCallable `json:"-"`
}

func (entity *UploadStickerFile) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *UploadStickerFile) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.PngSticker.(type) {
	case types.InputBytes:
		files["png_sticker"] = entity.PngSticker
		entity.PngSticker = nil
	}
	return files
}

func (UploadStickerFile) MethodName() string {
	return "uploadStickerFile"
}

func (UploadStickerFile) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.File `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeFile,
		Result: x1.Result,
	}
	return &result, nil
}
