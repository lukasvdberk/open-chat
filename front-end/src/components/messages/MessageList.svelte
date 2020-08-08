<script>
    import Message from "./Message.svelte";
    export let messages
</script>

{#each messages as message}
    <Message
        username="lukas yes yes"
        profilePhoto="https://avatars1.githubusercontent.com/u/38686669?s=128&u=94e13f84dc9e796a9d3a0485d90472f0fd4481b0&v=4"
        messageContent={message.messageContent}
    />
{/each}