// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResultCachedDocument Represents a link to a file stored on the Telegram servers
// By default, this file will be sent by the user with an optional caption
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// Note: This will only work in Telegram versions released after 9 April, 2016
// Older clients will ignore them.
type InlineQueryResultCachedDocument struct {
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         string                `json:"description,omitempty"`
	DocumentFileID      string                `json:"document_file_id"`
	ID                  string                `json:"id"`
	InputMessageContent any                   `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               string                `json:"title"`
}

func (entity InlineQueryResultCachedDocument) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                string                `json:"type"`
		ID                  string                `json:"id"`
		Title               string                `json:"title"`
		DocumentFileID      string                `json:"document_file_id"`
		Description         string                `json:"description,omitempty"`
		Caption             string                `json:"caption,omitempty"`
		ParseMode           string                `json:"parse_mode,omitempty"`
		CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent any                   `json:"input_message_content,omitempty"`
	}{
		Type:                "document",
		ID:                  entity.ID,
		Title:               entity.Title,
		DocumentFileID:      entity.DocumentFileID,
		Description:         entity.Description,
		Caption:             entity.Caption,
		ParseMode:           entity.ParseMode,
		CaptionEntities:     entity.CaptionEntities,
		ReplyMarkup:         entity.ReplyMarkup,
		InputMessageContent: entity.InputMessageContent,
	}
	if entity.InputMessageContent != nil {
		switch entity.InputMessageContent.(type) {
		case InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent, InputInvoiceMessageContent:
			break
		default:
			return nil, fmt.Errorf("input_message_content: unknown type: %T", entity.InputMessageContent)
		}
	}
	return json.Marshal(alias)
}
