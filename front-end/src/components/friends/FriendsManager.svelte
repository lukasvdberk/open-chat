<!-- Responsible for fetching and setting the data for the friends tab -->
<script>
    import {onMount} from "svelte";
    import {getAmountOfNewMessages, getFriends} from "./friends";
    import Friend from "./Friend.svelte";
    import {currentSelectedFriend, allCurrentFriends} from "./friends-store";

    let friends = []

    // Currently only pushed by push notifications to increase the new message counter
    allCurrentFriends.subscribe((newFriends) => {
        friends = newFriends
    })

    onMount(async () => {
        let tmpFriends = await getFriends()
        let amountOfFriendsNewMessages = await getAmountOfNewMessages()

        if(amountOfFriendsNewMessages !== undefined) {
            tmpFriends.forEach((friend) => {
                friend.amountOfNewMessages = amountOfFriendsNewMessages[friend.id]

                if(!friend.amountOfNewMessages) {
                    friend.amountOfNewMessages = 0
                }
            })
            friends = tmpFriends
            allCurrentFriends.set(friends)
        }
    })

    function onFriendSelected(event) {
        const friend = event.detail.friend

        // up date the store so other components can react to it

        if(friend.amountOfNewMessages >= 0) {
            friends.forEach((friendFromArr) => {
                if(friend.id === friendFromArr.id) {
                    friend.isActive = true
                    friend.amountOfNewMessages = 0
                    friendFromArr.isActive = true
                    friendFromArr.amountOfNewMessages = 0
                }
            })
        }
        currentSelectedFriend.set(friend)
    }
</script>

<style>
    section {
        width: 100%;
        overflow-y: auto;
    }
</style>
<section>
    <!-- List of friends   -->
    {#each friends as friend}
        <Friend
            id={friend.id}
            username={friend.username}
            profilePhoto={friend.profilePhoto}
            isActive={friend.isActive}
            amountOfNewMessages={friend.amountOfNewMessages}
            on:friend-select={onFriendSelected}
        />
    {/each}
</section>