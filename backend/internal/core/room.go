package core

import (
	"log"
	"sync"
	"time"
)

const ROOM_MAX_USERS int = 2

// Contains all server side game logic and game data
type Room struct {
	id         string
	players    map[string]*Player
	mutex      sync.RWMutex
	broadcast  chan []byte
	register   chan *Player
	unregister chan *Player
	tickRate   time.Duration
}

func (room *Room) Run() {
	ticker := time.NewTicker(room.tickRate)
	defer ticker.Stop()

	for {
		select {
		case player := <-room.register:
			room.registerPlayer(player)

		case player := <-room.unregister:
			room.unregisterPlayer(player)

		case player := <-room.broadcast:
			room.broadcastToAll(player)
		}
	}
}

func (room *Room) broadcastToAll(message []byte) {
	room.mutex.Lock()
	defer room.mutex.RUnlock()

	for _, player := range room.players {
		select {
		case player.sendChan <- message:
		default:
			log.Printf("Player %s send buffer is full!\n", player.id)
		}
	}
}

func (room *Room) unregisterPlayer(player *Player) {
	room.mutex.Lock()
	if _, exists := room.players[player.id]; exists {
		delete(room.players, player.id)
		close(player.sendChan)
	}
	room.mutex.Unlock()

	log.Printf("Player %s left the Room\n", player.id)
	room.sendGameState()
}

func (room *Room) registerPlayer(player *Player) {
	room.mutex.Lock()
	room.players[player.id] = player
	room.mutex.Unlock()

	log.Printf("Player %s joined the Room %s\n", player.id, room.id)
	room.sendGameState()
}

// Send all the current game state to
// all the connected users in the room
func (room *Room) sendGameState() {
	room.mutex.Lock()
	// Update dat
	room.mutex.Unlock()
}
