<script>
    import Icon from "../common/img/Icon.svelte";
    export let username;
    export let profilePhoto;
    export let isActive = false;
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

<div class:isActive>
    {#if profilePhoto !== ""}
        <Icon src={profilePhoto} alt={username} />
    {:else}
        <!--   TODO replace stock image with custom own made image     -->
        <Icon src="https://file.coffee/u/4SmZXSKoA.png" alt={username} />
    {/if}
    <span class:isActive>{username}</span>
</div>