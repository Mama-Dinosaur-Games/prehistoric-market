package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("Recieved: %s\n", msg)
		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("Write Error: ", err)
			break
		}
	}
}
