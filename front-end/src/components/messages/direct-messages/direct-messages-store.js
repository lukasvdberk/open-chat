// TODO add
// the friends store for the current active user.
import {writable} from "svelte/store";

// Strucucture like the following:
// [
//      friend-id: [
//            {
//                 "messageId": 36,
//                 "fromUser": 27,
//                 "messageContent": "hi",
//                 "readMessage": true,
//                 "sentAt": 0
//            },
//            {
//                 "messageId": 36,
//                 "fromUser": 28,
//                 "messageContent": "howdy",
//                 "readMessage": true,
//                 "sentAt": 123789123789
//            },
//      ]
//]
//
export const directMessages = writable([]);

export function saveMessageToStore(friendUserId, messageContent) {
    // maybe refactor to separate file or to direct-messages-store.js
    directMessages.update((currentMessages) => {
        console.log(currentMessages)
        if(currentMessages !== undefined) {
            // if it does not exist yet then we create it
            if(currentMessages[friendUserId] === undefined) {
                // create and add message object
                let newMessages = currentMessages
                newMessages[friendUserId] = [messageContent]
                return newMessages
            }
            // here we add it to existing messages
            else {
                // adds it to existing messages
                // TODO refactor with push
                currentMessages[friendUserId] = [...currentMessages[friendUserId], messageContent]
                return currentMessages
            }
        }
    })
}

export function saveMessagesToStore(friendUserId, messages) {
    directMessages.update((currentMessages) => {
        if(currentMessages !== undefined) {
            // if it does not exist yet then we create it
            if(currentMessages[friendUserId] === undefined) {
                let newMessages = currentMessages
                newMessages[friendUserId] = messages
                return newMessages
            }
            // here we add it to existing messages
            else {
                currentMessages[friendUserId] = currentMessages.concat(messages)
                return currentMessages
            }
        }
    })
}