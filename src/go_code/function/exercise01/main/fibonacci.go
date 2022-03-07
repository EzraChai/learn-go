package main

import "fmt"

func main() {
	// 1,1,2,3,5,8,13,21 -> Fibonacci Number
	fmt.Println(getFibonancci(4))
	fmt.Println(getFibonancci(5))
}

func getFibonancci(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return getFibonancci(num-1) + getFibonancci(num-2)
	}
}
