<script>
    // TODO probably need a event listener so the parent container can get what is being typed
    function onMessageSend(e) {
    }
</script>

<style>
    input {
        width: 100%;
        border-radius: var(--rounding);
    }
</style>

<form on:submit|preventDefault="{onMessageSend}">
    <input placeholder="type here to send a message" />
</form>
