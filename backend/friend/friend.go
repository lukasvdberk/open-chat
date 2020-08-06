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

func IsUserFriend(userId1 int64, userId2 int64) bool {
	// TODO implement for security reasons in messages.
	return true
}
