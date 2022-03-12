package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var arr [10]int

	insertRandomNumberIntoArray(&arr)
	bubbleSort(&arr)

	fmt.Println(arr)

	binarySearch(&arr, 90)

}

func insertRandomNumberIntoArray(arr *[10]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
	}
}

func bubbleSort(arr *[10]int) {
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1-j; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func binarySearch(arr *[10]int, findNum int) {
	leftIndex := 0
	rightIndex := len(arr) - 1
	for {
		if leftIndex > rightIndex {
			fmt.Println("Cant find number", findNum)
			return
		}
		middle := (leftIndex + rightIndex) / 2
		if arr[middle] > findNum {
			rightIndex = middle - 1
		} else if arr[middle] < findNum {
			leftIndex = middle + 1
		} else {
			fmt.Println(findNum, " found in index", middle)
			return
		}
	}
}
