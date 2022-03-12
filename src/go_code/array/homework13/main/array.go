package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var arr [5]int

	insertRandomNumberIntoArray(&arr)
	fmt.Println(arr)
	smallestNumIndex, biggestNumIndex := findSmallestNumAndBiggestNum(&arr)
	fmt.Println(arr[smallestNumIndex], ":", smallestNumIndex, " | ", arr[biggestNumIndex], ":", biggestNumIndex)
}

func insertRandomNumberIntoArray(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
	}
}

func findSmallestNumAndBiggestNum(arr *[5]int) (smallestNumIndex, biggestNumIndex int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[biggestNumIndex] {
			biggestNumIndex = i
		} else {
			if arr[i] < arr[smallestNumIndex] || i == 0 {
				smallestNumIndex = i
			}
		}
	}
	return
}
