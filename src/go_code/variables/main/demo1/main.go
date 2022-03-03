package main

import "fmt"

var (
	PI = 3.142
	e  = 2.71828
)

func main() {
	//Initialize variable
	var i int = 100

	var love string = "Chloe Gan"

	fmt.Println("i =", i)
	fmt.Println("I love", love)

	//	1 type
	var num int

	fmt.Println("num =", num)

	//	2 type
	var num2 = 123.431

	fmt.Println("num2 =", num2)

	//	3 type
	//	: is a must while initialize a variable in Go
	num3 := 234

	fmt.Println("num3 =", num3)

	//	Init a local variable in a line
	var n1, n2, n3 int
	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3)

	name1, name2 := "Chloe Gan", "Ezra Chai"
	fmt.Println("name1:", name1, "name2:", name2)

	fmt.Println("PI =", PI, "e =", e)
}
