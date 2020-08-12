// Listens when a service worker receives a notification. This will be trigged if the application is active.
// If you want to add a notification listener add it to the following list of notification listeners.


import {handleDirectMessageNotification} from "./notification-listeners/direct-messages";

export function setupNotificationListeners() {
    // The keys are the message channels that will bed added by the server to identify the type of notification
    const listeners = {
        "direct-messages": handleDirectMessageNotification,
    }

    const channel = new BroadcastChannel('sw-messages');
    channel.addEventListener('message', event => {
        const data = event.data.data
        const messageChannel = data.messageChannel
        const listener = listeners[messageChannel]

        if(listener !== undefined) {
            listener(data.data)
        }
    });
}