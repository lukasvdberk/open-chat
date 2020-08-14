package friend

import (
	"github.com/lukasvdberk/opensource-discord/auth"
	"github.com/lukasvdberk/opensource-discord/database"
	"strconv"
)

type Friend struct {
	Id    int64     `json:"id"`
	User1 auth.User `json:"user1"`
	User2 auth.User `json:"user2"`
}

func GetFriendsByUserId(userId int64) []auth.User {
	// Only the following data will be put in the user struct
	// id, username, profilephoto
	friendsFromUserListMap := database.SelectStatement(
		"SELECT User.id, User.username, User.profilePhoto  FROM User  JOIN Friend ON User.id = Friend.user1 OR User.id = Friend.user2   WHERE (Friend.user1 = ? OR Friend.user2 = ?) AND User.id != ?",
		userId, userId, userId,
	)

	if friendsFromUserListMap == nil {
		return nil
	}

	var friends []auth.User

	for _, friendMap := range friendsFromUserListMap {
		friend := new(auth.User)

		id := friendMap["id"]
		if _, err := strconv.Atoi(id); err == nil {
			friend.Id, _ = strconv.ParseInt(id, 10, 64)
		}

		friend.Username = friendMap["username"]
		friend.ProfilePhoto = friendMap["profilePhoto"]

		friends = append(friends, *friend)
	}

	return friends
}

func GetFriendRelation(userId1 int64, userId2 int64) int64 {
	friendRelationResults := database.SelectStatement(
		"SELECT id FROM Friend WHERE (user1 = ? AND user2 = ?) OR (user2 = ? AND user1 = ?)",
		userId1, userId2, userId1, userId2,
	)

	if len(friendRelationResults) > 0 {
		friendRelationId := friendRelationResults[0]["id"]
		if _, err := strconv.Atoi(friendRelationId); err == nil {
			actualId, _ := strconv.ParseInt(friendRelationId, 10, 64)
			return actualId
		}
	}

	// -1 means not found or a error occurred.
	return -1
}

func GetFriendsRelationsFromUserId(userId int64) []int64 {
	friendRelationResultsListMap := database.SelectStatement(
		"SELECT id FROM Friend WHERE user1 = ? OR user2 = ?",
		userId, userId,
	)

	var friendRelationIds []int64
	for _, friendRelationMap := range friendRelationResultsListMap {
		idStr := friendRelationMap["id"]
		if _, err := strconv.Atoi(idStr); err == nil {
			actualId, _ := strconv.ParseInt(idStr, 10, 64)
			friendRelationIds = append(friendRelationIds, actualId)
		}
	}

	return friendRelationIds
}
