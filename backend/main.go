package main

import (
	"backend/internal/core"
	"backend/internal/market"
	"fmt"
	"net/http"
)

func main() {
	_ = market.Market{}
	_ = core.CoreServer{}
	fmt.Println("Hello, World!")

	http.HandleFunc("/ws", core.HandleConnections)
	fmt.Println("WebSocket Started on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
