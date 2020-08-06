<script>
    import FriendsManager from "./components/friends/FriendsManager.svelte";
    import LoginOrRegister from "./components/auth/LoginOrRegister.svelte";
    import {onAuthenticationStateChange} from "./components/auth/auth-store";
    import {isAuthenticated} from "./components/auth/auth";

    let authenticated = isAuthenticated()

    // So it will redirect if the user is not logged in
    onAuthenticationStateChange.subscribe((value => {
        authenticated = value
    }))
</script>

<style>
    div {
        max-width: 15%;
    }
</style>

{#if !authenticated}
    <LoginOrRegister />
{:else}
    <main>
        <!--  TODO replace with container manager  -->
        <div>
            <FriendsManager />
        </div>
    </main>
{/if}