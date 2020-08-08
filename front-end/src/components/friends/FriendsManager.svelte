<!-- Responsible for fetching and setting the data for the friends tab -->
<script>
    import {onMount} from "svelte";
    import {getFriends} from "./friends";
    import Friend from "./Friend.svelte";
    import {currentSelectedFriend} from "./friends-store";


    let friends = []
    onMount(async () => {
        friends = await getFriends()
    })

    function onFriendSelected(event) {
        const friend = event.detail.friend

        // up date the store so other components can react to it
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
            id="{friend.id}"
            username={friend.username}
            profilePhoto={friend.profilePhoto}
            isActive={friend.isActive}
            on:friend-select={onFriendSelected}
        />
    {/each}
</section>