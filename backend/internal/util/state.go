package util

type ConnectionState int

// States can only change to the adjacent state, not any further
// InMainMenu <-> InLobby <-> InGame
const (
	InMainMenu ConnectionState = iota
	InLobby
	InGame
)
