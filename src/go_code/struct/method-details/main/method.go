package main

import "fmt"

type integer int

type Student struct {
	Name string
	Age  int
}

func (student *Student) String() string {
	return fmt.Sprintf("Name=[%v] Age=[%v]", student.Name, student.Age)
}

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

	stu := Student{
		Name: "Chloe Gan",
		Age:  0,
	}
	//	If you implement the String() function, fmt.Println() will use the String() function
	fmt.Println(&stu)
}
