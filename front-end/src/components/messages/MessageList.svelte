<script>
    import Message from "./Message.svelte";
    import {getUsername, getUserId} from "../auth/auth";
    import {currentSelectedFriend} from "../friends/friends-store";
    export let messages

    const thisClientUsername = getUsername()
    const thisClientUserId = getUserId()

    let selectedFriend = undefined
    console.log("current selected user",currentSelectedFriend)
    function getUsernameById(id) {
        // gets the username from this client
        if (thisClientUserId == id) {
            return thisClientUsername
        } else {
            if(selectedFriend !== undefined) {
                return selectedFriend.username
            }
        }
    }

    function getProfilePhoto(id) {
        if (thisClientUserId == id) {
            return "https://file.coffee/u/4SmZXSKoA.png"
        }
        return "https://avatars1.githubusercontent.com/u/38686669?s=128&u=94e13f84dc9e796a9d3a0485d90472f0fd4481b0&v=4"
    }

    currentSelectedFriend.subscribe((newSelectedFriend => {
        selectedFriend = newSelectedFriend
    }))
</script>

{#each messages as message}
    <Message
        username={getUsernameById(message.fromUser)}
        profilePhoto={getProfilePhoto(message.fromUser)}
        messageContent={message.messageContent}
    />
{/each}