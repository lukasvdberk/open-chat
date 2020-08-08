<script>
    import ProfileIcon from "../common/img/ProfileIcon.svelte";

    export let username = ''
    export let profilePhoto = ''
    export let messageContent = ''
    export let timeStamp = ''
    export let hasRead = ''
    // TODO add some prop to set whether it is this user or not
</script>

<style>
    div.container {
        margin: var(--rounding);
        width: auto;
        height: auto;
        padding: calc(var(--rounding)/2);
        background-color: var(--second-bg);
        border-radius: var(--rounding);
    }

    div.block {
        display:inline-block;
        padding: 0;
        margin: 0;
    }

    p {
        margin: 0;
        padding: 0;
    }
</style>

<div class="container">
    <div class="block">
        <ProfileIcon src={profilePhoto} alt={username} />
    </div>
    <div class="block">
        <b>{username}</b>
        <p>{messageContent}</p>
    </div>
</div>
