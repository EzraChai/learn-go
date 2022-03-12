package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 10)
	rand.Seed(time.Now().Unix())

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	//	Reverse an array
	arr = reverseArr(arr)
	fmt.Println(arr)

	number, index := getHighestNumber(arr)
	fmt.Println("Highest number =", number, "index =", index)

	flag := isThereNumber55(arr)
	fmt.Println("Is there any number 55?", flag)

}

//	Reverse an array
//1,2,3,4,5
//5,4,3,2,1
func reverseArr(arr []int) []int {
	length := len(arr) - 1
	for i := 0; i < (length/2)+1; i++ {
		//	Flip the place of a number
		arr[i], arr[length-i] = arr[length-i], arr[i]
	}
	return arr
}

func getHighestNumber(arr []int) (higestNumber, index int) {
	higestNumber = arr[0]
	for i := 1; i < len(arr); i++ {
		if higestNumber < arr[i] {
			higestNumber = arr[i]
			index = i
		}
	}
	return
}

func isThereNumber55(arr []int) (b bool) {
	leftIndex := 0
	rightIndex := len(arr) - 1
	searchNumber := 55
	for {
		if leftIndex > rightIndex {
			return
		}
		middleIndex := (leftIndex + rightIndex) / 2
		if arr[middleIndex] > searchNumber {
			rightIndex = middleIndex - 1
		} else if arr[middleIndex] < searchNumber {
			leftIndex = middleIndex + 1
		} else {
			return true
		}
	}
}
