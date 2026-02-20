import type { NetworkHandler } from "../net/NetworkHandler";

let instance: NetworkHandler | null = null;

export function network(): NetworkHandler {
    return instance!;
}

export function setNetworkHandler(app: NetworkHandler) {
    instance = app;
}
