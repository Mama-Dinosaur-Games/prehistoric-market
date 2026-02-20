package main

import (
	"backend/internal/core"
	"backend/internal/market"
	"fmt"
)

func main() {
	_ = market.Market{}
	_ = core.CoreServer{}
	fmt.Println("Hello, World!")
}
