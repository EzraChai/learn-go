package main

import "fmt"

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Stu struct {
	Name string
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()")
}
func (m Monster) Say() {
	fmt.Println("Monster Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

func (stu Stu) Say() {
	fmt.Println("Student Say()")
}

func main() {
	var stu = Stu{"Chloe Gan"}
	stu.Say()
	var I integer
	I = 12
	I.Say()

	m := Monster{}
	m.Hello()
	m.Say()
}
