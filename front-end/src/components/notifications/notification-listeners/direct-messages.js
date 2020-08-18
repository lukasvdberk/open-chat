import {addMessageIfChannelExits} from "../../messages/direct-messages/direct-messages-store";
import {allCurrentFriends, currentSelectedFriend} from "../../friends/friends-store";

let currentFriend = undefined
currentSelectedFriend.subscribe((newFriend) => {
    currentFriend = newFriend
})

export function handleDirectMessageNotification(data) {
    const userId = data.fromUser

    // If the current channel has no messages yet dont push the current message

    addMessageIfChannelExits(userId, data)
    allCurrentFriends.update((currentFriends) => {
        currentFriends.forEach((friend) => {
            if(friend.id === userId) {
                if(currentFriend !== undefined) {
                    if(currentFriend.id === userId) {
                        friend.isActive = true
                        return
                    }
                }

                if(friend.amountOfNewMessages) {
                    friend.amountOfNewMessages += 1
                }
                else {
                    friend.amountOfNewMessages = 1
                }
            }
            if(currentFriend !== undefined) {
                friend.isActive = currentFriend.id === friend.id
            }
        })
        return currentFriends
    })
}