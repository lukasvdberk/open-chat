
self.addEventListener("push", e => {
    clients.matchAll({
        type: 'window',
        includeUncontrolled: true
    }).then(function(windowClients) {
        // This push the data to a existing window or when none available show a notification
        const data = e.data.json();
        const channel = new BroadcastChannel('sw-messages');
        let clientIsVisible = false;

        for (let i = 0; i < windowClients.length; i++) {
            const windowClient = windowClients[i];
            if (windowClient.visibilityState==="visible") {
                clientIsVisible = true;
                channel.postMessage({
                    data: data
                });
            }
        }

        if(!clientIsVisible) {
            const popUpData = data.popUpData
            self.registration.showNotification(
                popUpData.title, {
                    body: popUpData.messageContent,
                    icon: popUpData.icon,
                }
            );
        }
    });
});