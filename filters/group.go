package filters

import "github.com/Squirrel-Network/gobotapi/types"

func Group() FilterOperand {
	return func(values ...any) bool {
		for _, value := range values {
			if chat, ok := value.(*types.Chat); ok && chat != nil {
				return chat.Type == "group"
			}
			if message, ok := value.(*types.Message); ok && message != nil {
				return message.Chat.Type == "group"
			}
		}
		return false
	}
}
