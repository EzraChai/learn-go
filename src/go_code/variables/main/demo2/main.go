package main

import "fmt"

//	Variables
func main() {
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i =", i)

	//Cannot change the type of the variables
	// i = 12.42 
	// fmt.Println("i =", i)


	//Cannot reintialize the variable with the same name
	// var i int = 50;

	var decimal float32 = 75634.3243
	fmt.Println(decimal)

}
