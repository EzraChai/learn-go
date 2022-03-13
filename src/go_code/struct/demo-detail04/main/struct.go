package main

import "fmt"

type A struct {
	Num int
}

type B struct {
	Num int
}

type C A

func main() {

	var a A
	var b B

	//a = b

	//	If you really want to migrate a different type of struct, make sure that the field and type are the same
	//	BUt why??!!
	a = A(b)

	fmt.Printf("a: %v b: %v\n", a, b)

	var c C
	//	Possible
	c = C(a)
	fmt.Println(c)

}
