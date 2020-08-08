<script>
    import Icon from "../common/img/Icon.svelte";
    export let username = ''
    export let profilePhoto = ''
    export let messageContent = ''
    export let timeStamp = ''
    export let hasRead = ''
    // TODO probably refactor with something like small icon
</script>

<style>
    div {
        padding: var(--rounding);
        margin: var(--rounding);
        width: auto;
        height: auto;

        background-color: var(--second-bg);
        border-radius: var(--rounding);
    }

    span {

    }
</style>
<div>
    <Icon src={profilePhoto} alt={username} />
    <span>{messageContent}</span>
</div>
