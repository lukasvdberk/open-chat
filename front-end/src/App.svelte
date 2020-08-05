<script>
    import FriendsManager from "./components/friends/FriendsManager.svelte";
    import LoginOrRegister from "./components/auth/LoginOrRegister.svelte";
    import Login from "./components/auth/Login.svelte";
    import Register from "./components/auth/Register.svelte";
    import {onAuthenticationStateChange} from "./components/auth/auth-store";
    import {isAuthenticated} from "./components/auth/auth";

    let authenticated = isAuthenticated()
    onAuthenticationStateChange.subscribe((value => {
        console.log(authenticated)
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