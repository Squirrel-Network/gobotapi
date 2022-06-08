// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// GetMyDefaultAdministratorRights Use this method to get the current default administrator rights of the bot
// Returns ChatAdministratorRights on success.
type GetMyDefaultAdministratorRights struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

func (entity *GetMyDefaultAdministratorRights) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *GetMyDefaultAdministratorRights) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetMyDefaultAdministratorRights) MethodName() string {
	return "getMyDefaultAdministratorRights"
}

func (GetMyDefaultAdministratorRights) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.ChatAdministratorRights `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeChatAdministratorRights,
		Result: x1.Result,
	}
	return &result, nil
}
