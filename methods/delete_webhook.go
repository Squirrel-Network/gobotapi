// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// DeleteWebhook Use this method to remove webhook integration if you decide to switch back to getUpdates
// Returns True on success.
type DeleteWebhook struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

func (entity *DeleteWebhook) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *DeleteWebhook) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (DeleteWebhook) MethodName() string {
	return "deleteWebhook"
}

func (DeleteWebhook) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result bool `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
