package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var intArr [5]int
	for i, _ := range intArr {
		intArr[i] = rand.Intn(100)
	}

	fmt.Println(intArr)
	flippedArray := flipArr(intArr)
	fmt.Println(flippedArray)
}

func flipArr(intArray [5]int) (arr [5]int) {
	//length := len(intArray) - 1
	//for i := 0; i <= length; i++ {
	//	arr[length-i] = intArray[i]
	//}
	temp := 0
	length := len(intArray)
	for i := 0; i < length/2; i++ {
		temp = intArray[length-1-i]
		intArray[length-1-i] = intArray[i]
		intArray[i] = temp
	}
	return intArray
}
