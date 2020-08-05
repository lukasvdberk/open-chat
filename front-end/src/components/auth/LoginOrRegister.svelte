<script>
    import Login from "./Login.svelte";
    import Register from "./Register.svelte";

    let showLogin = false
    let showRegister = false
</script>

{#if !showLogin && !showRegister}
    <div>
        <button on:click="{() => {showLogin = true}}">Login</button>
        <button on:click="{() => {showRegister = true}}">Register</button>
    </div>
{/if}

{#if showLogin}
    <Login />
{/if}

{#if showRegister}
    <Register />
{/if}