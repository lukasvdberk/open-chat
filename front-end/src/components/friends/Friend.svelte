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
    {#if profilePhoto !== ""}
        <img src={profilePhoto} alt={username} />
    {:else}
        <!--   TODO replace stock image with custom own made image     -->
        <img src="https://file.coffee/u/4SmZXSKoA.png" alt={username} />
    {/if}
    <span class:isActive>{username}</span>
</div>