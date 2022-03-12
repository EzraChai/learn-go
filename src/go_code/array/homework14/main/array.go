package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var arr [8]int

	insertRandomNumberIntoArray(&arr)
	fmt.Println(arr)

	avgNum := getAverageNum(&arr)
	fmt.Println("Average Number:", avgNum)

	numHigherThanAverage, numLowerThanAverage := getNumHigerThanAverageAndNumLowerThanAverage(&arr, avgNum)
	fmt.Println("Num higher than average:", numLowerThanAverage, " Num higher than average:", numHigherThanAverage)

}

func insertRandomNumberIntoArray(arr *[8]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
	}
}

func getAverageNum(arr *[8]int) (avgNum float64) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr))
}

func getNumHigerThanAverageAndNumLowerThanAverage(arr *[8]int, averageNum float64) (numHigherThanAverage, numLowerThanAverage []int) {
	for i := 0; i < len(arr); i++ {
		if float64(arr[i]) > averageNum {
			numHigherThanAverage = append(numHigherThanAverage, arr[i])
		} else if float64(arr[i]) < averageNum {
			numLowerThanAverage = append(numLowerThanAverage, arr[i])
		}
	}
	return
}
