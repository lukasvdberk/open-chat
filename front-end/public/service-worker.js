
self.addEventListener("push", e => {
    const data = e.data.json();


    const channel = new BroadcastChannel('sw-messages');
    // TODO check if active then either push a notification or push the data to a existing tab.
    const hasActiveTab = false
    clients.matchAll({includeUncontrolled: true, type: 'window'}).then((result) => {
        console.log(result)
        channel.postMessage({
            data: data
        });
    })

    self.registration.showNotification(data.title, {
        body: data.messageContent,
    });
});