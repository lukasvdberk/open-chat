<script>
    import Icon from "../common/img/Icon.svelte";
    import {createEventDispatcher} from "svelte";
    import {currentSelectedFriend} from "./friends-store";
    import ProfileIcon from "../common/img/ProfileIcon.svelte";

    export let id;
    export let username;
    export let profilePhoto;
    export let isActive = false;

    let dispatch = createEventDispatcher();

    function onFriendSelected() {
        dispatch("friend-select", {
            friend: {
                id: id,
                username: username,
                profilePhoto: profilePhoto,
            }
        });
    }

    currentSelectedFriend.subscribe((newSelectedFriend => {
        isActive = false
        if (newSelectedFriend) {
            if (newSelectedFriend.id === id) {
                isActive = true
            }
        }
    }))
</script>

<style>
    div {
        position: relative;
        margin: 8px;
        border-radius: var(--rounding);
        background-color: var(--second-bg);
    }

    div:hover {
        cursor: pointer;
    }

    .isActive {
        color: var(--opposite-text);
        background-color: var(--is-active) !important;
    }

    span {
        font-size: 22px;
        vertical-align: top;
    }
</style>

<div class:isActive on:click={onFriendSelected}>
    <ProfileIcon src={profilePhoto} alt={username} />
    <span class:isActive>{username}</span>
</div>