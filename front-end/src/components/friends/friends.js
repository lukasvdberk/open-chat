// Responsible for fetching friends of a user.

import {get} from "../api/request";

export async function getFriends() {
    const response = await get("friend")

    if (response.code !== 0) {
        return []
    }
    return response.content.friends
}

export async function getAmountOfNewMessages() {
    const response = await get("amount_of_new_messages")

    if (response.code !== 0) {
        return undefined
    }
    return response.content.amountOfMessagesPerUserId
}