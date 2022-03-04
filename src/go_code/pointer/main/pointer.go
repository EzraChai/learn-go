package main

import "fmt"

func main() {

	var i int = 10
	// i's address
	fmt.Println("i's address", &i)

	//	1. ptr is a pointer variable
	//	2. ptr type is *int
	//	3. ptr value is &i (0xc00001e0e0)
	var ptr *int = &i
	fmt.Println("value:", ptr)
	fmt.Println("address:", &ptr)
	fmt.Println("dereference:", *ptr)
}
