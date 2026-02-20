package core

import (
	"net"
	"sync"
)

type CoreServer struct {
	listener net.Listener
	mu       sync.RWMutex
	quit     chan struct{}
}
