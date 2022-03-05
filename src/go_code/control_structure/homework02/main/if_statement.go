package main

import "fmt"

func main() {
	var num1, num2 float64
	num1 = 12.4
	num2 = 4.3

	if num1 > 10.0 && num2 < 20.0 {
		fmt.Println(num1 + num2)
	}
}
