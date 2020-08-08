<script>
    // TODO add handler for friend direct message on change
    // TODO add handler for when a user is visiting a channel of a server

    // TODO save received messages in stores
    import MessageList from "./MessageList.svelte";
    import SendMessage from "./SendMessage.svelte";
    import {getFriendMessages} from "./direct-messages";
    import {onMount} from "svelte";

    let messages = []
    onMount(async () => {
        const response = await getFriendMessages(30)
        if (response !== undefined) {
            console.log("fetching messages")
            messages = response.messages
        } else {
            console.log("failed to retrieve messages")
        }
    })
    // TODO replace with actual data
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
