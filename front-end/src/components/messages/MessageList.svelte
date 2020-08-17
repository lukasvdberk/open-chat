<script>
    import Message from "./Message.svelte";
    import {getUsername, getUserId} from "../auth/auth";
    import {currentSelectedFriend} from "../friends/friends-store";
    import {afterUpdate, createEventDispatcher, onMount} from "svelte";

    export let messages

    const thisClientUsername = getUsername()
    const thisClientUserId = getUserId()

    let selectedFriend = undefined

    let dispatch = createEventDispatcher();

    function onReachedTop(ignored) {
        dispatch("reached-top")
    }


    function getUsernameById(id) {
        // gets the username from this client
        if (thisClientUserId == id) {
            return thisClientUsername
        } else {
            if(selectedFriend !== undefined) {
                return selectedFriend.username
            }
            return 'Could not get username'
        }
    }

    function getProfilePhoto(id) {
        // TODO refactor this method bc this is not what it is.
        if (thisClientUserId == id) {
            return "https://file.coffee/u/4SmZXSKoA.png"
        }
        return "https://avatars1.githubusercontent.com/u/38686669?s=128&u=94e13f84dc9e796a9d3a0485d90472f0fd4481b0&v=4"
    }

    currentSelectedFriend.subscribe((newSelectedFriend => {
        selectedFriend = newSelectedFriend
    }))

    function getMessageContainer() {
        return document.getElementById("message-container")
    }

    function scrollToBottom() {
        let messageContainer = getMessageContainer()

        if(document.getElementById("message-container")) {
            // so it always scroll to the bottom
            messageContainer.scrollTo({
                left: 0,
                top: messageContainer.scrollHeight,
                behavior: "smooth"
            });
        }
    }

    function onScroll(event) {
        if(event.target.scrollTop === 0) {
            // then we reached the top of the page and need to fetch new messages.
            onReachedTop()
        }
    }

    afterUpdate(() => {
        scrollToBottom()
    })
</script>

<style>
    section {
        overflow-y: scroll;
        max-height: 100%;
    }
</style>

<section id="message-container" on:scroll={onScroll}>
    {#if messages.length !== 0}
        {#each messages as message}
            <Message
                    username={getUsernameById(message.fromUser)}
                    profilePhoto={getProfilePhoto(message.fromUser)}
                    messageContent={message.messageContent}
                    sentAt={message.sentAt}
                    hasNotRead={!message.readMessage}
            />
        {/each}
    {:else}
        <!--Make this message more beautiful-->
        <p style="color: var(--opposite-text)">No messages yet. Be the first one to start a conversation by sending a message below</p>
    {/if}
</section>
