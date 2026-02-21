package main

import (
	"backend/internal/core"
	"backend/internal/market"
)

func main() {
	_ = market.Market{}
	server := core.CoreServer{}
	server.Init()
}
