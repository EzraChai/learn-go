package main

import "fmt"

func test(intPtr *int) {
	*intPtr += 10
}

func main() {
	n := 10
	test(&n)

	fmt.Println(n) // 20
}
