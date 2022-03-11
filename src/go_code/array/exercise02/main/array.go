package main

import "fmt"

func main() {
	var numArr01 [3]int = [3]int{1, 2, 3}
	for _, i2 := range numArr01 {
		fmt.Printf("%v ", i2)
	}
	fmt.Println()

	var numArr02 = [3]int{4, 5, 6}
	fmt.Println(numArr02)

	var numArr03 = [...]int{8, 9, 10}
	fmt.Println(numArr03)
}
