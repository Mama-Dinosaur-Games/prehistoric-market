package market

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Market struct {
	priceIndex map[string]float32
}

func (m *Market) getStockPrice(stockName string) float32 {
	return m.priceIndex[stockName]
}

type Stock struct {
	Name       string
	Price      float64
	Volatility float64
}

func runSimulation(days <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	day := 1

	// Initialize everything
	stocks := []Stock{
		{"MEAT", 100.0, 0.005},
		{"FISH", 50.0, 0.01},
		{"PLANT", 20.0, 0.002},
	}

	for range days {

		fmt.Printf("Day: %d\n", day) // placeholder for main functionality
		for i := range stocks {
			stocks[i].UpdatePrice(day)
			fmt.Printf("%s: %.2f\n", stocks[i].Name, stocks[i].Price)
		}

		// Do more updates daily here

		day++
	}
}

func (s *Stock) UpdatePrice(day int) {
	increase := 0.005 // increase volatility by 0.5% per day
	vol := s.Volatility + increase*float64(day)

	epsilon := rand.NormFloat64() * vol
	s.Price *= 1 + epsilon
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
