package main

import "fmt"

//	It's like class in JAVA
func AddUpper() func(int) int {
	var n int = 10           //	declare variable
	return func(x int) int { //	a function in a class
		n += x
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1)) //	11
	fmt.Println(f(2)) //	13
	fmt.Println(f(3)) //	16
}
