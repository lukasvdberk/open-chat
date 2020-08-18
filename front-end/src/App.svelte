<script>
    import FriendsManager from "./components/friends/FriendsManager.svelte";
    import LoginOrRegister from "./components/auth/LoginOrRegister.svelte";
    import {onAuthenticationStateChange} from "./components/auth/auth-store";
    import {isAuthenticated} from "./components/auth/auth";
    import MessageManager from "./components/messages/MessageManager.svelte";
    import {registerDevice} from "./components/notifications/device";
    import {setupNotificationListeners} from "./components/notifications/notification-listener";
    import {setUserInfoFromApi} from "./components/auth/user-info";

    let authenticated = isAuthenticated()

    // So it will redirect if the user is not logged in
    onAuthenticationStateChange.subscribe((value => {
        // register service worker. currently only needed for push notifications
        if(value) {
            setUserInfoFromApi()
            registerDevice().then()
            setupNotificationListeners()
        }
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