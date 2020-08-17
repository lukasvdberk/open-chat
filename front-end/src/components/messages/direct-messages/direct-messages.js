import {get, post} from "../../api/request";

export async function getFriendMessages(userId) {
    const response = await get(`messages/${userId}/-1`)

    if(response.code === 0) {
        return response.content
    }
    return undefined
}

export async function getFriendMessagesFromTimestamp(userId, fromTimestamp) {
    const response = await get(`messages/${userId}/${fromTimestamp}`)

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

export async function updateReadMessages(friendUserId) {
    const response = await get(`read_messages/${friendUserId}`)

    return response.code === 0;
}