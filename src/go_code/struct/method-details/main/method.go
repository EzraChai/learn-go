package main

import "fmt"

type integer int

func (i *integer) print() {
	fmt.Println("i =", *i)
}

func (i *integer) numberAdd1() {
	*i++
}

func main() {
	var i integer = 10
	i.print()
	i.numberAdd1()
	i.print()
}
