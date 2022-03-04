package main

import "fmt"

func main() {

	//	init a variable
	var num int = 920
	//	get the num variable's pointer
	fmt.Println(&num)

	//	init a variable pointing the num variable's pointer
	var ptr *int = &num

	//	use the pointer to change the num variable's value
	*ptr = 25

	//	print the num variable
	fmt.Println(num)
}
