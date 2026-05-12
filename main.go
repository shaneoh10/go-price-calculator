package main

import (
	"github.com/shaneoh10/go-price-calculator/cmdmanager"
	// "github.com/shaneoh10/go-price-calculator/filemanager"
	"github.com/shaneoh10/go-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.09, 0.12, 0.21}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", taxRate*100))
		cmdm := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		job.Process()
	}
}
