import Alpine from 'alpinejs';

export function startPolling() {
    console.log("Polling...");

    setTimeout(startPolling, 1000);
}

Alpine.start()
