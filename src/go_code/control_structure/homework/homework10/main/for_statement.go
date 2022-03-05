package main

import "fmt"

func main() {

	var sum int

	for i := 9; i <= 100; i += 9 {
		fmt.Println(i)
		sum += i
	}
	fmt.Println("SUM =", sum)

}
