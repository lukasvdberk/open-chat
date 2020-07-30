// Responsible for fetching friends of a user.

export async function getFriends() {
    // TODO replace with actual api fetching

    let friends = []
    for (let i = 0; i < 30; i++) {
        friends.push({
            username: "Lukas",
            profilePhoto: "https://avatars1.githubusercontent.com/u/38686669?s=64&u=94e13f84dc9e796a9d3a0485d90472f0fd4481b0&v=4"
        })
    }

    return friends
}