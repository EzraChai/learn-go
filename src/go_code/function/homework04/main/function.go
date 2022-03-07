package main

import "fmt"

func pyramid(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {

	var n int
	fmt.Println("Please enter a number")
	fmt.Scanln(&n)
	pyramid(n)

}
