<script>
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

    section.messages {
        display: block;
        position: absolute;
        top: 40%;
        left: 85%;
        height: 25px;
        width: 25px;
        border-radius: 25px;
        background-color: var(--attention-color);
        z-index: 2;
        text-align: center;
    }

    section > span {
        color: var(--opposite-text);
        font-size: 25px;
    }
</style>

<div class:isActive on:click={onFriendSelected}>
    <ProfileIcon src={profilePhoto} alt={username} />
    <span class:isActive>{username}</span>
    <section class="messages">
        <span>5</span>
    </section>
</div>