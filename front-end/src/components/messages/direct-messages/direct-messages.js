import {get, post} from "../../api/request";

export async function getFriendMessages(userId) {
    const response = await get(`messages/${userId}`)

    if(response.code === 0) {
        return response.content
    }
    return undefined
}

export async function sendFriendMessage(friendUserId, messageStr) {
    const response = await post(`messages/`, {
        friendUserId: friendUserId,
        messageContent: messageStr
    })

    if(response.code === 0) {
        return response.content
    }
    return undefined
}
