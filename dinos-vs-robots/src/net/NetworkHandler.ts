export class NetworkHandler {
    private webSocket: WebSocket
    constructor(url: URL = new URL("ws://localhost:8080/ws")) {
        this.webSocket = new WebSocket(url)
    }

    async initSocket() {
        this.webSocket.onopen = function(event) {
            console.log(event.type)
            console.log("Connected To WebSocket!\n");
        };

        this.webSocket.onmessage = function(event) {
            console.log(event.type)
            console.log("Recieved: " + event.data + "\n");
        };

        this.webSocket.onclose = function(event) {
            console.log(event.type)
            console.log("Disconnected from WebSocket!\n");
        };
    }

    sendMsg() {
        const message: string = "Hello!"
        this.webSocket.send(message);
    }
}
