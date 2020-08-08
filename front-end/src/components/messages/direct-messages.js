import {get} from "../api/request";

async function getFriendMessages(userId) {
    const response = await get(`messages/${userId}`)

    if(response.code === 0) {
        return response.content
    }
    return undefined
}

async function saveFriendMessage(userId) {

}
