package util

import "net"

type Connection struct {
	addr  net.Addr
	state ConnectionState
}
