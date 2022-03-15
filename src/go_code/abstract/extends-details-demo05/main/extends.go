package main

import "fmt"

type C struct {
	Name int
}

type A struct {
	Name string
	age  int
	C
}

func (a *A) SayOk() {
	fmt.Println("A SayOk")
}

func (a *A) hello() {
	fmt.Println("A Hello,", a.Name)
}

func (b *B) hello() {
	fmt.Println("B Hello,", b.Name)
}

type B struct {
	A
}

func main() {

	var b B
	b.Name = "Chloe Gan"
	b.age = 20
	b.SayOk()
	b.A.hello()
	b.hello()
}
