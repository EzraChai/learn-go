package main

import "fmt"

func main() {
	var result int
	fmt.Println("Enter your result")
	fmt.Scanln(&result)

	if result < 100 && result >= 0 {
		switch {
		case result >= 60:
			fmt.Println("Pass")
		default:
			fmt.Println("Fail")

		}
	}
}
