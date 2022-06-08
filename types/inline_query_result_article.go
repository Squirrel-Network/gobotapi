// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResultArticle Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	Description         string                `json:"description,omitempty"`
	HideURL             bool                  `json:"hide_url,omitempty"`
	ID                  string                `json:"id"`
	InputMessageContent any                   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int64                 `json:"thumb_width,omitempty"`
	Title               string                `json:"title"`
	URL                 string                `json:"url,omitempty"`
}

func (entity InlineQueryResultArticle) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                string                `json:"type"`
		ID                  string                `json:"id"`
		Title               string                `json:"title"`
		InputMessageContent any                   `json:"input_message_content"`
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		URL                 string                `json:"url,omitempty"`
		HideURL             bool                  `json:"hide_url,omitempty"`
		Description         string                `json:"description,omitempty"`
		ThumbURL            string                `json:"thumb_url,omitempty"`
		ThumbWidth          int64                 `json:"thumb_width,omitempty"`
		ThumbHeight         int                   `json:"thumb_height,omitempty"`
	}{
		Type:                "article",
		ID:                  entity.ID,
		Title:               entity.Title,
		InputMessageContent: entity.InputMessageContent,
		ReplyMarkup:         entity.ReplyMarkup,
		URL:                 entity.URL,
		HideURL:             entity.HideURL,
		Description:         entity.Description,
		ThumbURL:            entity.ThumbURL,
		ThumbWidth:          entity.ThumbWidth,
		ThumbHeight:         entity.ThumbHeight,
	}
	switch entity.InputMessageContent.(type) {
	case InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent, InputInvoiceMessageContent:
		break
	default:
		return nil, fmt.Errorf("input_message_content: unknown type: %T", entity.InputMessageContent)
	}
	return json.Marshal(alias)
}
