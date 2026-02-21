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
	Points     int
}


type Portfolio struct {
	Player int
	Shares map[string]int
}

func runSimulation(days <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	day := 1

	// Initialize everything
	stocks := []Stock{
		{"MEAT", 100.0, 0.005, 1},
		{"FISH", 50.0, 0.01, 1},
		{"PLANT", 20.0, 0.002, 1},
		{"EGG", 0.00, 0.00, 1},
		{"INSECT", 0.00, 0.00, 1},
	}

	portfolios := []Portfolio{
		{0, map[string]int{"MEAT": 0, "FISH": 0, "PLANT": 0, "EGG": 0, "INSECT": 0}},
		{1, map[string]int{"MEAT": 0, "FISH": 0, "PLANT": 0, "EGG": 0, "INSECT": 0}},
	}

	for range days {

		fmt.Printf("Day: %d\n", day) // placeholder for main functionality
		for i := range stocks {
			stocks[i].UpdatePrice(day)
			fmt.Printf("%s: %.2f\n", stocks[i].Name, stocks[i].Price)
		}

		for _, p := range portfolios {
			fmt.Printf("Player %d:\n", p.Player)
			for stock, qty := range p.Shares {
				fmt.Printf("  %s: %d shares\n", stock, qty)
			}
		}

		// Do more updates daily here

		day++
	}
}

func (s *Stock) UpdatePrice(day int) {
	increase := 0.005 // increase volatility by 0.5% per day
	vol := s.Volatility + increase*float64(day)

	epsilon := rand.NormFloat64() * vol

	// Update price based off the points
	pointFactor := 0.002
	pointEffect := float64(s.Points) * pointFactor

	s.Price *= 1 + epsilon + pointEffect

}	

func (s *Portfolio) buyStock(portfolios []Portfolio, player int, stockName string){
	portfolios[player].Shares[stockName] ++
}

func (s *Portfolio) sellStock(portfolios []Portfolio, player int, stockName string){
	portfolios[player].Shares[stockName] --
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
