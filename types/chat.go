package types

type Chat struct {
	Bio string `json:"bio,omitempty"`
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
	Description string `json:"description,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	HasPrivateForwards bool `json:"has_private_forwards,omitempty"`
	HasProtectedContent bool `json:"has_protected_content,omitempty"`
	Id int64 `json:"id"`
	InviteLink string `json:"invite_link,omitempty"`
	LastName string `json:"last_name,omitempty"`
	LinkedChatId int64 `json:"linked_chat_id,omitempty"`
	Location *ChatLocation `json:"location,omitempty"`
	MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`
	Permissions *ChatPermissions `json:"permissions,omitempty"`
	Photo *ChatPhoto `json:"photo,omitempty"`
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`
	StickerSetName string `json:"sticker_set_name,omitempty"`
	Title string `json:"title,omitempty"`
	Type string `json:"type"`
	Username string `json:"username,omitempty"`
}