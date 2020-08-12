import {get, post} from "../api/request";

function urlBase64ToUint8Array(base64String) {
    const padding = "=".repeat((4 - base64String.length % 4) % 4);
    const base64 = (base64String + padding)
        .replace(/\-/g, "+")
        .replace(/_/g, "/");

    const rawData = window.atob(base64);
    const outputArray = new Uint8Array(rawData.length);

    for (let i = 0; i < rawData.length; ++i) {
        outputArray[i] = rawData.charCodeAt(i);
    }
    return outputArray;
}

export async function registerDevice() {
    // TODO check if already registered
    // Register Service Worker
    const scopeUrl = "/"
    const registration = await navigator.serviceWorker.getRegistration(scopeUrl);

    if (registration === undefined) {
        // If no active service worker is found we want to register it.
        const register = await navigator.serviceWorker.register("/service-worker.js", {
            scope: scopeUrl
        });

        // Make sure it registered before continuing.
        await navigator.serviceWorker.ready;

        // Getting the public push key from the server
        const response = await get("web-notifications/get_public_key")
        const publicKey = response.content.publicKey

        // Register Push
        const subscription = await register.pushManager.subscribe({
            userVisibleOnly: true,
            applicationServerKey: urlBase64ToUint8Array(publicKey)
        });

        await post("web-notifications/add-device", subscription)
    } else {
        // This will fetch the new service-worker.js script.
        await registration.update()
    }
}