// TODO add
// the friends store for the current active user.
import {writable} from "svelte/store";

// Should set the following data
// id of ther user
// username
// profilePhoto
// Set all that in a Object
export const currentSelectedFriend = writable(undefined);