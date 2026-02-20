package market

import (
	"fmt"
	"sync"
	"time"
)

type Market struct {
	priceIndex map[string]float32
}

func (m *Market) getStockPrice(stockName string) float32 {
	return m.priceIndex[stockName]
}

func runSimulation(days <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	day := 1

	for range days {

		fmt.Printf("Day: %d\n", day) // placeholder for main functionality
		day++
	}
}

func Run() {
	nextDay := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)

	go runSimulation(nextDay, &wg)

	// run fake simulation of days
	time.Sleep(1 * time.Second)
	nextDay <- true

	time.Sleep(1 * time.Second)
	nextDay <- true

	time.Sleep(1 * time.Second)
	nextDay <- true

	time.Sleep(1 * time.Second)
	nextDay <- true

	close(nextDay)
	wg.Wait()

	fmt.Println("Hello, World!")
}
