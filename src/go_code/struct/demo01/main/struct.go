package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	var cat1 Cat
	cat1.Name = "Hello Kitty"
	cat1.Age = 10
	cat1.Color = "white"
	fmt.Println(cat1)
	fmt.Println("name =", cat1.Name)
	fmt.Println("age =", cat1.Age)
	fmt.Println("color =", cat1.Color)
}
