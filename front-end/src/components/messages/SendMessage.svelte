<script>
    import {createEventDispatcher} from "svelte";

    let dispatch = createEventDispatcher();

    let messageContent = ''
    function onMessageSend(ignored) {
        dispatch("message-sent", {messageContent: messageContent})
        messageContent = ""
    }

</script>

<style>
    input {
        width: 100%;
        border-radius: var(--rounding);
    }
</style>


<form on:submit|preventDefault="{onMessageSend}">
    <input bind:value={messageContent} placeholder="type here to send a message" />
</form>
