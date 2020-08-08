<script>
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
    <img src={src} alt={alt} />
</div>