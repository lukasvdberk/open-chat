<script>
    // TODO add handler for when a user is visiting a channel of a server

    // TODO save received messages in stores
    import MessageList from "./MessageList.svelte";
    import SendMessage from "./SendMessage.svelte";
    import {getFriendMessages} from "./direct-messages";
    import {currentSelectedFriend} from "../friends/friends-store";

    let messages = []

    async function setMessages(currentFriend) {
        if(currentFriend !== undefined) {
            // TODO maybe also store messages in svelte stores.
            const response = await getFriendMessages(currentFriend.id)
            if (response !== undefined) {
                messages = response.messages
            } else {
                console.log("failed to retrieve messages")
            }
        }
    }

    currentSelectedFriend.subscribe((newSelectedFriend => {
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
    <MessageList messages={messages} />
</div>
<div class="send-message">
    <SendMessage />
</div>
