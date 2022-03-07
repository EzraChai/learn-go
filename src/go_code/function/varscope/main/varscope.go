package main

import "fmt"

var age int = 16
var Name string = "Chloe Gan"

func test() {
	//	The variables' (age and Name) scope is only in this function
	//age := 16
	//Name := "Chloe Gan"
}

func main() {
	//fmt.Println(Name)	//	undefined: Name
	fmt.Println(Name) //	Chloe Gan
	fmt.Println(age)  //	16

	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
	}
}
