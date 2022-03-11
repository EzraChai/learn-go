package main

import "fmt"

func main() {

	var numArr = []int{1, -1, 9, 90, 12}
	sum, average := getSumAndAverage(numArr)
	fmt.Println("Sum:", sum, "average:", average)
}

func getSumAndAverage(arr []int) (sum int, averageNum float64) {
	for _, value := range arr {
		sum += value
	}
	averageNum = float64(sum) / float64(len(arr))
	return
}
