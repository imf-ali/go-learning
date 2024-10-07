package main

import (
	"fmt"

	// "imf-ali.com/first-app/cmdmanager"
	"imf-ali.com/first-app/filemanager"
	"imf-ali.com/first-app/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		// cmdm := cmdmanager.New()
		priceJob := prices.New(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
      fmt.Println("Failed to process job")
			fmt.Println(err)
    } 
	}
}