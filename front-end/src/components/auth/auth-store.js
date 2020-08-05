// TODO make a svelte store here were in app.svelte is subscribed to so if the user logs out it will redirect to login screen.

import {writable} from 'svelte/store';
import {isAuthenticated} from "./auth";

export const onAuthenticationStateChange = writable(isAuthenticated());