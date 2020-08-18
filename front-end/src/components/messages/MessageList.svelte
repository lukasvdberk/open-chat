<script>
    import Message from "./Message.svelte";
    import {getUsername, getUserId} from "../auth/auth";
    import {allCurrentFriends, currentSelectedFriend} from "../friends/friends-store";
    import {afterUpdate, beforeUpdate, createEventDispatcher, onMount} from "svelte";
    import {getUserInfo} from "../auth/user-info";

    export let messages

    const thisClientUsername = getUsername()
    const thisClientUserId = getUserId()

    let selectedFriend = undefined
    let noNewMessages = false

    let friends = []
    allCurrentFriends.subscribe((newFriends) => {
        friends = newFriends
    })

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
            return getUserInfo().profilePhoto
        }

        for(let i = 0; i < friends.length; i++) {
            let friend = friends[i]
            if(friend.id == id) {
                return friend.profilePhoto
            }
        }
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

    let oldScrollHeight = 0
    function onScroll(event) {
        if(event.target.scrollTop === 0) {
            // then we reached the top of the page and need to fetch new messages.
            onReachedTop()
        }
    }

    let oldMessageSize = 0
    afterUpdate(() => {
        const currentLength = messages.length

        if(oldMessageSize === 0) {
            scrollToBottom()
        }

        // If there was 1 message added we want to show it by scrolling to the bottom.
        if(oldMessageSize + 1 === currentLength) {
            scrollToBottom()
        }

        oldMessageSize = messages.length

        if(getMessageContainer() !== undefined) {
            oldScrollHeight =  getMessageContainer().scrollTop
        }
    })
</script>

<style>
    section {
        overflow-y: scroll;
        max-height: 100%;
    }

    p {
        color: var(--opposite-text)
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
        <!-- Make this message more beautiful -->
        <p style="color: var(--opposite-text)">No messages yet. Be the first one to start a conversation by sending a message below</p>
    {/if}
</section>
