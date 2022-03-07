package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println(n) // 2 2 3
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println(n) // 2
	}
}

func main() {
	//test(4)

	test2(4)
}
