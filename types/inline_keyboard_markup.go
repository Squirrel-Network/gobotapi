// Code AutoGenerated; DO NOT EDIT.

package types

// InlineKeyboardMarkup Represents an inline keyboard that appears right next to the message it belongs to.
// Note: This will only work in Telegram versions released after 9 April, 2016
// Older clients will display unsupported message.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}
