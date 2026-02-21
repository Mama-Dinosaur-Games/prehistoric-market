package core

import (
	"backend/internal/util"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
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

type CoreServer struct {
	connections map[net.Addr]*util.Connection
	rooms       map[int]*Room // lobbies get unique ids
	roomCount   int
	listener    net.Listener
	mu          sync.RWMutex
	quit        chan struct{}
}

func (server *CoreServer) Init() {
	// Default values
	server.roomCount = 0

	http.HandleFunc("/ws", handleConnections)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

func (server *CoreServer) createLobby(id string) *Room {
	room := &Room{
		id:         id,
		players:    make(map[string]*Player),
		mutex:      sync.RWMutex{},
		broadcast:  make(chan []byte, 256),
		register:   make(chan *Player),
		unregister: make(chan *Player),
		tickRate:   time.Millisecond * 50, // 20 Updates per second
	}

	go room.Run()
	return room
}

func (server *CoreServer) getLobbies() {

}
