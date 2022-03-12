package main

import (
	"fmt"
)

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := arr[:]
	slice = append(slice, 11)
	copyArr := make([]int, 11)
	copy(copyArr, slice)

	for i := 0; i < len(copyArr); i++ {
		fmt.Printf("%v ", copyArr[i])
	}
}
