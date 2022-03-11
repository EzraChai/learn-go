package main

import "fmt"

//	array
func main() {
	var hens [6]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 10.0

	totalWeight := 0.0
	for _, hen := range hens {
		totalWeight += hen
	}
	fmt.Println("Total Weight :", totalWeight)
	avgWeight := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Println("Avg Weight :", avgWeight)
}
