package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println(n) // 2 2 3
}

func main() {
	test(4)
}
