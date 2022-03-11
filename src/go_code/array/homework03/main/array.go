package main

import "fmt"

func main() {

	var numArr = [...]int{1, 2, 2, 4, 5, 5, 6, 6, 7, 7, 9}
	sum, average := getSumAndAverage(numArr)
	fmt.Println("Sum:", sum, "average:", average)
}

func getSumAndAverage(arr [11]int) (sum int, averageNum float64) {
	for _, value := range arr {
		sum += value
	}
	averageNum = float64(sum / len(arr))
	return
}
