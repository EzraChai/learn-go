package main

import "fmt"

var (
	Func1 = func(n1, n2 int) int {
		return n1 + n2
	}
)

func main() {

	res := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)

	a := func(n1, n2 int) int {
		return n1 + n2
	}

	fmt.Println(res)
	fmt.Println(a(10, 30))

	res4 := Func1(70, 40)
	fmt.Println(res4)
}
