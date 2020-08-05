<script>
    export let username;
    export let profilePhoto;
    export let isActive = false;
</script>

<style>
    div {
        position: relative;
        margin: 8px;
        border-radius: 8px;
        background-color: var(--second-bg);
    }

    img, div {
        height: 48px;
        vertical-align: middle;
    }

    .isActive {
        color: var(--opposite-text);
        background-color: var(--is-active) !important;
    }
</style>

<div class:isActive>
    <img src={profilePhoto} alt={username} />
    <span class:isActive>{username}</span>
</div>