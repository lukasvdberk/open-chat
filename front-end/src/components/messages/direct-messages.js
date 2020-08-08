import {get} from "../api/request";

export async function getFriendMessages(userId) {
    const response = await get(`messages/${userId}`)

    if(response.code === 0) {
        return response.content
    }
    return undefined
}

export async function saveFriendMessage(userId) {

}
