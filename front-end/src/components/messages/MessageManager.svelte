<script>
    // TODO add handler for friend direct message on change
    // TODO add handler for when a user is visiting a channel of a server

    // TODO save received messages in stores
    import MessageList from "./MessageList.svelte";
    import SendMessage from "./SendMessage.svelte";

    // TODO replace with actual data
    let messages = [
        {
            "id": 14,
            "fromUser": 14,
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596886503
        },
        {
            "id": 15,
            "fromUser": 15,
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596886510
        },
        {
            "id": 16,
            "fromUser": 16,
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596887152
        },
        {
            "id": 17,
            "fromUser": 17,
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596893510
        },
        {
            "id": 20,
            "fromUser": 20,
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596899118
        },
        {
            "id": 21,
            "fromUser": 21,
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596899127
        },
        {
            "id": 22,
            "fromUser": 22,
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596899128
        }
    ]
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
