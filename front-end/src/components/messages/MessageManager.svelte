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
            "fromUser": {
                "id": 14,
                "username": "",
                "profilePhoto": ""
            },
            "friendRelation": {
                "id": 3,
                "user1": {
                    "id": 0,
                    "username": "",
                    "profilePhoto": ""
                },
                "user2": {
                    "id": 0,
                    "username": "",
                    "profilePhoto": ""
                }
            },
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596886503
        },
        {
            "id": 15,
            "fromUser": {
                "id": 15,
                "username": "",
                "profilePhoto": ""
            },
            "friendRelation": {
                "id": 3,
                "user1": {
                    "id": 0,
                    "username": "",
                    "profilePhoto": ""
                },
                "user2": {
                    "id": 0,
                    "username": "",
                    "profilePhoto": ""
                }
            },
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596886510
        },
        {
            "id": 16,
            "fromUser": {
                "id": 16,
                "username": "",
                "profilePhoto": ""
            },
            "friendRelation": {
                "id": 3,
                "user1": {
                    "id": 0,
                    "username": "",
                    "profilePhoto": ""
                },
                "user2": {
                    "id": 0,
                    "username": "",
                    "profilePhoto": ""
                }
            },
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596887152
        },
        {
            "id": 17,
            "fromUser": {
                "id": 17,
                "username": "",
                "profilePhoto": ""
            },
            "friendRelation": {
                "id": 3,
                "user1": {
                    "id": 0,
                    "username": "",
                    "profilePhoto": ""
                },
                "user2": {
                    "id": 0,
                    "username": "",
                    "profilePhoto": ""
                }
            },
            "messageContent": "dikke paarden stront",
            "readMessage": false,
            "sentAt": 1596893510
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
