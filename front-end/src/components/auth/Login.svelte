<script>
    import {login} from "./auth";

    let username = ''
    let password = ''

    function loginUser() {
        login(username, password).then((result) => {
            if (result) {
                alert("Username and password do not exist in the database.")
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

<button on:click={loginUser}>Login</button>