<script>
    import Message from "./Message.svelte";
    import {getUsername, getUserId} from "../auth/auth";
    export let messages

    const thisClientUsername = getUsername()
    const thisClientUserId = getUserId()

    function getUsernameById(id) {
        // gets the username from this client
        if (thisClientUserId == id) {
            return thisClientUsername
        } else {
            // TODO fetch other username from friends store or something.
        }
    }
</script>

{#each messages as message}
    <Message
        username={getUsernameById(message.fromUser)}
        profilePhoto="https://avatars1.githubusercontent.com/u/38686669?s=128&u=94e13f84dc9e796a9d3a0485d90472f0fd4481b0&v=4"
        messageContent={message.messageContent}
    />
{/each}