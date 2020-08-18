The front-end of the open source discord written in svelte.


###Notes

####Css variables
All colours or general information use must be defined in css variables. Set the variables in public/vars.css


####Requests
The base url of the api resides in api/request.js. Change that if you have the api server running at a different location
Also only make request through the api/request.js methods.

###Notifications
Notifications are either shown as notification (duh) or when a browser tab is active pushed to the current tab.

To add you own notification handler (for example new messages or something like that.) Do the following steps:
1. Create a js file in src/notifications/notification-listeners/yourlistener.js
2. Then export a function for example. In this file you can everything you want. I would recommend communicating with a svelte store of some sort and not edit the document directly.
```js
export function handleDirectMessageNotification(data) {
    // TODO implement details
}
```
3. Register that in function in notification listener.js in the listeners variable like this
```js
"direct-messages": handleDirectMessageNotification
```
The key (direct-messages in this case) tells what kind of notification this is. So when a notification is received from the backend it will set a messagechannel which will have key like that to tel what kind of notification it is.

