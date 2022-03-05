package main

import "fmt"

func main() {
	var num1, num2 int32
	num1 = 5
	num2 = 10
	sum := num1 + num2
	if sum%3 == 0 && sum%5 == 0 {
		fmt.Println("Nice!!")
	}
}
