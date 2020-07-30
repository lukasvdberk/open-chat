<!-- Responsible for fetching and setting the data for the friends tab -->
<script>
    import {onMount} from "svelte";
    import {getFriends} from "./friends";
    import Friend from "./Friend.svelte";


    let friends = []
    onMount(async () => {
        friends = await getFriends()
        friends[0].isActive = true
    })
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
            username={friend.username}
            profilePhoto={friend.profilePhoto}
            isActive={friend.isActive}
        />
    {/each}
</section>