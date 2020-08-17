<script>
    import ProfileIcon from "../common/img/ProfileIcon.svelte";

    export let username = ''
    export let profilePhoto = ''
    export let messageContent = ''
    export let hasNotRead = false
    // is a unix timestamp
    export let sentAt = 0

    // TODO add some prop to set whether it is this user or not

    function getReadableDateTimeFromTimestamp(timestamp) {
        let a = new Date(timestamp * 1000);
        let months = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'];
        let year = a.getFullYear();
        let month = months[a.getMonth()];
        let date = a.getDate();
        let hour = a.getHours();
        let min = a.getMinutes();
        let sec = a.getSeconds();
        return  + hour + ':' + min + ':' + sec + ' ' + date + ' ' + month + ' ' + year;
    }

    // After a couple of seconds just say he read it
    if(hasNotRead) {
        // TODO make this a fun animation
        setTimeout(() => hasNotRead = false, 2000)
    }
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

    .hasNotRead {
        background-color: var(--attention-color) !important;
    }

    div.img-block {
        width: 30px !important;
        vertical-align: top;
        height: auto;
    }

    div.block {
        display:inline-block;
        padding: 0;
        margin: 0;
        width: calc(100% - 40px);
    }

    div.block > * {
        display: block;
    }

    div.img-block > * {
        vertical-align: text-top;
    }

    p {
        margin: 0;
        padding: 0;
    }

    i, b {
        display: inline-block;
        float: left;
    }
</style>

<div class:hasNotRead class="container">
    <div class="block img-block">
        <ProfileIcon src={profilePhoto} alt={username} />
    </div>
    <div class="block">
        <b>{username}</b>
        <i>   - received at: {getReadableDateTimeFromTimestamp(sentAt)}</i>
        <br>
        <p>{messageContent}</p>
    </div>
</div>
