package direct

import (
	"errors"
	"github.com/lukasvdberk/opensource-discord/auth"
	"github.com/lukasvdberk/opensource-discord/database"
	"github.com/lukasvdberk/opensource-discord/friend"
	"strconv"
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

func GetMessagesFromFriend(friendRelationId int64) []FriendMessage {
	messagesListMap := database.SelectStatement("SELECT * FROM FriendMessage WHERE friendRelation = ? ORDER BY sentAt",
		friendRelationId,
	)

	var messages []FriendMessage

	for _, messageMap := range messagesListMap {
		message := new(FriendMessage)

		id := messageMap["id"]
		if _, err := strconv.Atoi(id); err == nil {
			message.Id, _ = strconv.ParseInt(id, 10, 64)
		}

		fromUser := messageMap["id"]
		if _, err := strconv.Atoi(fromUser); err == nil {
			// maybe also add the other data such as username
			message.FromUser.Id, _ = strconv.ParseInt(fromUser, 10, 64)
		}

		message.FriendRelation.Id = friendRelationId
		message.ReadMessage = false
		message.SentAt = database.DateStringToTimeStamp(messageMap["sentAt"])
		messages = append(messages, *message)
	}

	return messages
}
