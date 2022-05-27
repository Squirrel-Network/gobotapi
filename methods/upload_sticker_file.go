package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type UploadStickerFile struct {
	PngSticker rawTypes.InputFile `json:"png_sticker,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *UploadStickerFile) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.PngSticker.(type) {
		case types.InputFile:
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
	result := rawTypes.Result {
		Kind: types.TypeFile,
		Result: x1.Result,
	}
	return &result, nil
}