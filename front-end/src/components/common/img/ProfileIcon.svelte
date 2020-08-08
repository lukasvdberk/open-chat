<script>
    import Icon from "./Icon.svelte";
    export let src
    export let alt
</script>

<style>
    img, div {
        padding: 0;
        margin: 0;
        display: inline-block;
        height: 30px;
        width: auto;
    }
</style>

<div>
    {#if src !== ""}
        <Icon src={src} alt={alt} />
    {:else}
    <!--   TODO replace stock image with custom own made image     -->
        <Icon src="https://file.coffee/u/4SmZXSKoA.png" alt={alt} />
    {/if}
</div>