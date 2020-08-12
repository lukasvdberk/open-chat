import {addMessageIfChannelExits} from "../../messages/direct-messages/direct-messages-store";

export function handleDirectMessageNotification(data) {
    const userId = data.fromUser

    // If the current channel has no messages yet dont push the current message

    addMessageIfChannelExits(userId, data)
}