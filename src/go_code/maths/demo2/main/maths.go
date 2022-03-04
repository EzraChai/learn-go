package main

import "fmt"

func main() {
	var i int = 0

	//a := i++
	i++

	//if i++ > 0 {
	if i > 0 {
		fmt.Println("ok")
	}

	a := 100
	pointer := &a
	fmt.Println(pointer, *pointer)
}
