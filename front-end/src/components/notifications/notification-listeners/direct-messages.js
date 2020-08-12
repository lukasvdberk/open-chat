import {saveMessageToStore} from "../../messages/direct-messages/direct-messages-store";

export function handleDirectMessageNotification(data) {
    saveMessageToStore(data.fromUser, data)
}