package main

import "fmt"

func main() {
	//	2D array
	/*
		0 0 0 0 0 0
		0 0 1 0 0 0
		0 2 0 3 0 0
		0 0 0 0 0 0
	*/
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	for _, ints := range arr {
		for _, value := range ints {
			fmt.Printf("%v ", value)
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}
