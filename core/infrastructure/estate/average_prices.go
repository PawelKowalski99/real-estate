package estate

import (
	"context"
	"net/http"
)

type AveragePrice struct {
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

func (t *Logic) GetAveragePrices(ctx context.Context) (avgPrices map[string]AveragePrice, err error, httpErr int) {
	pricesRows, err := t.queries.GetPrices(ctx)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	pricesMap := map[string]AveragePrice{}

	for _, row := range pricesRows {
		if row.Price > 0 {
			avgPrice := pricesMap[row.City]
			avgPrice.Price += row.Price
			avgPrice.Count++
			pricesMap[row.City] = avgPrice
		}
	}

	return pricesMap, nil, http.StatusOK
}

func (t *Logic) GetAveragePricesPerM2(ctx context.Context) (avgPrices map[string]AveragePrice, err error, httpErr int) {
	pricesRows, err := t.queries.GetPricesPerM2(ctx)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	pricesMap := map[string]AveragePrice{}

	for _, row := range pricesRows {
		if row.PricePerM2 > 0 {
			avgPrice := pricesMap[row.City]
			avgPrice.Price += row.PricePerM2
			avgPrice.Count++
			pricesMap[row.City] = avgPrice
		}
	}

	return pricesMap, nil, http.StatusOK
}
