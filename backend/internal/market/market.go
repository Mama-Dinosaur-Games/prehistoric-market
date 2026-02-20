package market

type Market struct {
	priceIndex map[string]float32
}

func (m *Market) getStockPrice(stockName string) float32 {
	return m.priceIndex[stockName]
}
