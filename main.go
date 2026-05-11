package main

import (
	"github.com/shaneoh10/go-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.09, 0.12, 0.21}

	for _, taxRate := range taxRates {
		job := prices.NewTaxIncludedPriceJob(taxRate)
		job.Process()
	}
}
