<script>
    // TODO add handler for when a user is visiting a channel of a server

    // TODO save received messages in stores
    import MessageList from "./MessageList.svelte";
    import SendMessage from "./SendMessage.svelte";
    import {getFriendMessages, sendFriendMessage} from "./direct-messages/direct-messages";
    import {currentSelectedFriend} from "../friends/friends-store";
    import {directMessages, saveMessagesToStore, saveMessageToStore} from "./direct-messages/direct-messages-store";

    // will be set by the store
    let allMessages = undefined

    // only messages from a specific channel.
    let channelMessages = []
    let currentFriend = undefined

    async function setMessages(currentFriend) {
        if(currentFriend) {
            // means they are already messages for this channel and we dont need to fetch.

            const channelMessagesFromStore = allMessages[currentFriend.id]
            if(Array.isArray(channelMessagesFromStore)) {
                channelMessages = channelMessagesFromStore
            }
            else {
                const response = await getFriendMessages(currentFriend.id)
                if (response !== undefined) {
                    let messages = response.messages

                    if(messages == null) {
                        messages = []
                    }
                    console.log(messages)
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

    directMessages.subscribe((newMessages) => {
        allMessages = newMessages
        setMessages(currentFriend)
    })

    // this will be triggered when user switches to another friend in the side pane
    currentSelectedFriend.subscribe((newSelectedFriend => {
        currentFriend = newSelectedFriend
        setMessages(newSelectedFriend).then
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
    <MessageList messages={channelMessages} />
</div>
<div class="send-message">
    <SendMessage on:message-sent={onMessageEnter}/>
</div>
