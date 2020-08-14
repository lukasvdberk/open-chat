<!-- Responsible for fetching and setting the data for the friends tab -->
<script>
    import {onMount} from "svelte";
    import {getAmountOfNewMessages, getFriends, updateReadMessages} from "./friends";
    import Friend from "./Friend.svelte";
    import {currentSelectedFriend} from "./friends-store";


    let friends = []
    let amountOfFriendsNewMessages = {}

    onMount(async () => {
        let tmpFriends = await getFriends()
        amountOfFriendsNewMessages = await getAmountOfNewMessages()

        if(amountOfFriendsNewMessages !== undefined) {
            tmpFriends.forEach((friend) => {
                friend.amountOfNewMessages = amountOfFriendsNewMessages[friend.id]

                if(!friend.amountOfNewMessages) {
                    friend.amountOfNewMessages = 0
                }
            })
            friends = tmpFriends
        }

    })

    function onFriendSelected(event) {
        const friend = event.detail.friend
        // up date the store so other components can react to it
        currentSelectedFriend.set(friend)

        if(amountOfFriendsNewMessages[friend.id]) {
            updateReadMessages(friend.id).then((result) => {
                console.log("sdkjflaksdjflkasdlkfjklj")
                friend.amountOfNewMessages = 0
                friends.splice(friend, 1)
                friend.isActive = true

                console.log(friend)
                friends = [...friends, friend]
            })
        }
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
            id="{friend.id}"
            username={friend.username}
            profilePhoto={friend.profilePhoto}
            isActive={friend.isActive}
            amountOfNewMessages={friend.amountOfNewMessages}
            on:friend-select={onFriendSelected}
        />
    {/each}
</section>