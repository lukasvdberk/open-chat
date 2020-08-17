<script>
    // TODO add handler for when a user is visiting a channel of a server

    // TODO save received messages in stores
    import MessageList from "./MessageList.svelte";
    import SendMessage from "./SendMessage.svelte";
    import {
        getFriendMessages,
        getFriendMessagesFromTimestamp,
        sendFriendMessage,
        updateReadMessages
    } from "./direct-messages/direct-messages";
    import {currentSelectedFriend} from "../friends/friends-store";
    import {directMessages, saveMessagesToStore, saveMessageToStore} from "./direct-messages/direct-messages-store";

    // will be set by the store
    let allMessages = undefined

    // only messages from a specific channel.
    let channelMessages = []
    let currentFriend = undefined

    function sortMessagesByTimestamp(messageList) {
        const sortByTimestamp = (first, second) => {
            if(first.sentAt > second.sentAt) return 1;
            if(second.sentAt > first.sentAt) return -1;
        }
        return messageList.sort(sortByTimestamp)
    }
    async function setMessages(currentFriend) {
        if(currentFriend) {
            // means they are already messages for this channel and we dont need to fetch.
            const channelMessagesFromStore = allMessages[currentFriend.id]

            if(Array.isArray(channelMessagesFromStore)) {
                channelMessages = sortMessagesByTimestamp(channelMessagesFromStore)
            }

            else {
                // If the messages are not there we need to fetch them
                const response = await getFriendMessages(currentFriend.id)
                if (response !== undefined) {
                    let messages = response.messages

                    if(messages == null) {
                        messages = []
                    }

                    saveMessagesToStore(currentFriend.id, messages)
                } else {
                    console.log("failed to retrieve messages")
                }
            }
        }
    }

    // when the user submits a new message to sent.
    function onMessageEnter(event) {
        const messageContent = event.detail.messageContent

        if(messageContent && currentSelectedFriend) {
            // will send it to the api
            sendFriendMessage(currentFriend.id, messageContent)
                .then((response) => {
                    const messageContent = response.messageContent
                    const friendUserId = currentFriend.id
                    saveMessageToStore(friendUserId, messageContent)
                })
        }
    }

    function onReachedTop() {
        // TODO when at the start of a conversation it keeps getting not the latest messages.
        // of the current channel
        let firstMessageTimestamp = (Math.min(...channelMessages.map(message => message.sentAt)))

        // fetch new messages
        getFriendMessagesFromTimestamp(currentFriend.id, firstMessageTimestamp).then((response) => {
            if (response !== undefined) {
                let messages = response.messages

                if(messages == null) {
                    messages = []
                }

                saveMessagesToStore(currentFriend.id, messages)
            } else {
                console.log("failed to retrieve messages")
            }
        })
    }

    directMessages.subscribe((newMessages) => {
        allMessages = newMessages

        if (currentFriend !== undefined) {
            setMessages(currentFriend)
        }
    })

    // this will be triggered when user switches to another friend in the side pane
    currentSelectedFriend.subscribe((newSelectedFriend => {
        currentFriend = newSelectedFriend
        setMessages(newSelectedFriend).then((result) => {
            // He has retrieved and opened the messages we now update that he messages have been read.
            updateReadMessages(currentFriend.id)
        })
    }))
</script>

<style>
    div {
        position: relative;
    }
    div.message-manager {
        height: 95%;
    }
    div.send-message {
        height: 5%;
    }
</style>

<div class="message-manager">
    <MessageList
        messages={channelMessages}
        on:reached-top={onReachedTop}
    />
</div>
<div class="send-message">
    <SendMessage on:message-sent={onMessageEnter}/>
</div>
