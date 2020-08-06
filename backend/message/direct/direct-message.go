package direct

import (
	"errors"
	"github.com/lukasvdberk/opensource-discord/auth"
	"github.com/lukasvdberk/opensource-discord/database"
	"github.com/lukasvdberk/opensource-discord/friend"
)

// this is the same as an dm
type FriendMessage struct {
	Id             int64         `json:"id"`
	FromUser       auth.User     `json:"fromUser"`
	FriendRelation friend.Friend `json:"friendRelation"`
	MessageContent string        `json:"messageContent"`
	ReadMessage    bool          `json:"readMessage"`
	// will be a timestamp
	SentAt int64 `json:"sentAt"`
	// TODO add file support
	//File    	   int64         `json:"sentAt"`
}

func SaveMessage(message FriendMessage) (FriendMessage, error) {
	// Only requires the following information
	// 		message.FromUser,
	//		message.FriendRelation,
	//		message.MessageContent,
	//		message.ReadMessage,
	//@return bool returns if message was successfully saved.

	messageId := database.InsertStatement("INSERT INTO FriendMessage VALUES(DEFAULT, ?, ?, ?, ?, NOW(), NULL);",
		message.FromUser.Id,
		message.FriendRelation.Id,
		message.MessageContent,
		message.ReadMessage,
	)

	if messageId != -1 {
		message.Id = messageId
		return message, nil
	} else {
		return message, errors.New("failed to save message to database")
	}
}
