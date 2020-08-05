<script>
    import {register} from "./auth";

    let username = ''
    let password = ''

    function registerUser() {
        register(username, password).then((result) => {
            if (result) {
               alert("Failed to register the user")
            }
        })
    }
</script>

<h1>Welcome, now you can login!</h1>
<label>
    Username
    <input placeholder="username" bind:value={username} />
</label>

<label>
    Password
    <input placeholder="password" type="password" bind:value={password} />
</label>

<button on:click={registerUser}>Register</button>