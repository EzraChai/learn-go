package main

import "fmt"

func main() {
	arr := [5]int{1, 3, 5, 7, 9}
	reverseArr(&arr)
	fmt.Println(arr)
}

//	Reverse an array
func reverseArr(arr *[5]int) *[5]int {

	for i := 0; i < len(arr)/2; i++ {
		arr[len(arr)-1-i], arr[i] = arr[i], arr[len(arr)-1-i]
	}
	return arr
}
