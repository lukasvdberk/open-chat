import {saveMessageToStore} from "../../messages/direct-messages/direct-messages-store";

export function handleDirectMessageNotification(data) {
    // TODO add to message store.
    console.log(data)
    console.log(data.fromUser)
    saveMessageToStore(data.fromUser, data)
}