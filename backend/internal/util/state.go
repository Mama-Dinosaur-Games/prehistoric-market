package util

type ConnectionType int

// States can only change to the adjacent state, not any further
// InMainMenu <-> InLobby <-> InGame
const (
	InMainMenu ConnectionType = iota
	InLobby
	InGame
)
