package main

import (
	"fmt"

	// "github.com/shaneoh10/go-price-calculator/cmdmanager"
	"github.com/shaneoh10/go-price-calculator/filemanager"
	"github.com/shaneoh10/go-price-calculator/prices"
)

const pricesFilePath = "prices.txt"

func main() {
	taxRates := []float64{0, 0.09, 0.12, 0.21}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New(pricesFilePath, fmt.Sprintf("result_%v.json", taxRate*100))
		// cmdm := cmdmanager.New()

		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go job.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
