<script>
    import FriendsManager from "./components/friends/FriendsManager.svelte";
    import LoginOrRegister from "./components/auth/LoginOrRegister.svelte";
    import {onAuthenticationStateChange} from "./components/auth/auth-store";
    import {isAuthenticated} from "./components/auth/auth";
    import MessageManager from "./components/messages/MessageManager.svelte";

    let authenticated = isAuthenticated()

    // So it will redirect if the user is not logged in
    onAuthenticationStateChange.subscribe((value => {
        authenticated = value
    }))
</script>

<style>
    main {
        display: block;
        height: 100vh;
    }

    div {
        height: 100%;
    }
    div.friends {
        width: 15%;
        float: left;
    }

    div.messages {
        width: 85%;
        float: right;
    }
</style>

{#if !authenticated}
    <LoginOrRegister />
{:else}
    <main>
        <div class="friends">
            <FriendsManager />
        </div>
        <div class="messages">
            <MessageManager />
        </div>
    </main>
{/if}