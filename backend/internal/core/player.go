package core

type PlayerActionType int

const ()

type PlayerAction struct {
}

type Player struct {
	id       string
	sendChan chan []byte
}
