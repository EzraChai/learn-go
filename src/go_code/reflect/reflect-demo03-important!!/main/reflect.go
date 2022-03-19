package main

import (
	"fmt"
	"reflect"
)

const (
	a = iota + 10
	_
	b
	c
	d = iota
)

const (
	e = iota
)

type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	rValue := reflect.ValueOf(b)
	fmt.Println(rValue)

	rKind := rValue.Kind()
	fmt.Println(rKind)

	value := rValue.Interface().(*Student)
	fmt.Println("Name:", value.Name)
}

func main() {
	fmt.Println(a, b, c, d, e)

	var stu *Student = &Student{
		Name: "Chloe Gan",
		Age:  16,
	}
	reflectTest02(stu)
}
