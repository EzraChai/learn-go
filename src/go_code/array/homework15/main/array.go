package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr [8]int

	insertRandomNumberIntoArray(&arr)
	fmt.Println(arr)

}

func insertRandomNumberIntoArray(arr *[8]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
	}
}
