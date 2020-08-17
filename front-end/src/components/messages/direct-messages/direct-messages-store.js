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
//                 "messageId": 37,
//                 "fromUser": 28,
//                 "messageContent": "howdy",
//                 "readMessage": true,
//                 "sentAt": 123789123789
//            },
//      ]
//]
//
export const directMessages = writable([]);

// Edge case for push notifications. When the channel is not filled with messages yet it will not fetch the latest messages.
// DONT USE THIS EXCEPT IN A NOTIFICATION HANDLER
export function addMessageIfChannelExits(friendUserId, messageContent) {
    directMessages.update((currentMessages) => {
        if(currentMessages !== undefined) {
            // if it does not exist yet then we create it
            if(currentMessages[friendUserId] !== undefined) {
                // adds it to existing messages
                currentMessages[friendUserId] = [...currentMessages[friendUserId], messageContent]
                return currentMessages
            }
        }
        return []
    })
}

export function saveMessageToStore(friendUserId, messageContent) {
    directMessages.update((currentMessages) => {
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
                currentMessages[friendUserId] = currentMessages[friendUserId].concat(messages)
                return currentMessages
            }
        }
    })
}