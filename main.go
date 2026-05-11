package main

import (
	"fmt"

	"github.com/shaneoh10/go-price-calculator/filemanager"
	"github.com/shaneoh10/go-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.09, 0.12, 0.21}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", taxRate*100))
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		job.Process()
	}
}
