<script>
    export let src
    export let alt
</script>

<style>
    img, div {
        display: inline-block;
        height: 48px;
        vertical-align: middle;
    }
</style>

<div>
    <img src={src} alt={alt} />
</div>