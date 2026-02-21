package core

import "backend/internal/util"

type Player struct {
	id       string
	sendChan chan []byte
}
