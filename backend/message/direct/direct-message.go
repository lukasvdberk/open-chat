package direct

import (
	"errors"
	"github.com/lukasvdberk/opensource-discord/database"
	"github.com/lukasvdberk/opensource-discord/friend"
	"strconv"
)

// this is the same as an dm
type FriendMessage struct {
	Id int64 `json:"messageId"`
	// Just the user id
	FromUser       int64         `json:"fromUser"`
	FriendRelation friend.Friend `json:"-"`
	MessageContent string        `json:"messageContent"`
	ReadMessage    bool          `json:"readMessage"`
	// will be a timestamp
	SentAt int64 `json:"sentAt"`
	// TODO add file support
	//File    	   int64         `json:"sentAt"`
}

func SaveMessage(message *FriendMessage) (*FriendMessage, error) {
	// Only requires the following information
	// 		message.FromUser,
	//		message.FriendRelation,
	//		message.MessageContent,
	//		message.ReadMessage,
	//@return bool returns if message was successfully saved.

	messageId := database.InsertStatement("INSERT INTO FriendMessage VALUES(DEFAULT, ?, ?, ?, ?, NOW(), NULL);",
		message.FromUser,
		message.FriendRelation.Id,
		message.MessageContent,
		message.ReadMessage,
	)

	if messageId != -1 {
		// TODO check whether you can directly fetch the created timestamp.
		// to get the timestamp from the database. else it may not be accurate

		// list of maps
		createdMessage := database.SelectStatement("SELECT sentAt FROM FriendMessage WHERE id=? LIMIT 1", messageId)
		if len(createdMessage) > 0 {
			message.SentAt = database.DateStringToTimeStamp(createdMessage[0]["sentAt"])
		}

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

		fromUser := messageMap["fromUser"]
		if _, err := strconv.Atoi(fromUser); err == nil {
			// TODO maybe also add the other data such as username
			message.FromUser, _ = strconv.ParseInt(fromUser, 10, 64)
		}

		message.FriendRelation.Id = friendRelationId

		// TODO set the right condition
		message.ReadMessage = messageMap["readMessage"] == "1"
		message.SentAt = database.DateStringToTimeStamp(messageMap["sentAt"])
		message.MessageContent = messageMap["messageContent"]

		messages = append(messages, *message)
	}

	return messages
}

// Returns a map of the friendId as key and how many new messages there are as value
func GetAmountOfNewMessagesFromUser(userId int64) map[int64]int64 {
	amountOfMessagesPerFriend := database.SelectStatement("SELECT Friend.id, COUNT(*) AS `amount_of_messages` "+
		"FROM Friend "+
		"JOIN FriendMessage ON Friend.id = FriendMessage.friendRelation "+
		"WHERE Friend.user1 = ? OR Friend.user2 = ? GROUP BY Friend.id",
		userId, userId,
	)

	mapToReturn := make(map[int64]int64)
	for _, friendAmountMessages := range amountOfMessagesPerFriend {
		var id int64 = -1
		idStr := friendAmountMessages["id"]
		if _, err := strconv.Atoi(idStr); err == nil {
			id, _ = strconv.ParseInt(idStr, 10, 64)
		}

		var amountOfMessages int64 = -1
		amountOfMessagesStr := friendAmountMessages["amount_of_messages"]
		if _, err := strconv.Atoi(amountOfMessagesStr); err == nil {
			amountOfMessages, _ = strconv.ParseInt(amountOfMessagesStr, 10, 64)
		}

		if amountOfMessages != -1 && id != -1 {
			mapToReturn[id] = amountOfMessages
		}
	}

	return mapToReturn
}
