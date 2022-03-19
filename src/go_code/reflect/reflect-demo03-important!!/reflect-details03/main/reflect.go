package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	var num int = 10
	reflectNum(&num)
	fmt.Println("i :", num)

	stu := Student{
		Name: "",
		Age:  17,
	}

	reflect01(&stu)
	fmt.Println()
	fmt.Println(stu)

}

func reflect01(b interface{}) {
	rValue := reflect.ValueOf(b)
	fmt.Println(rValue)

	iV := rValue.Elem().Interface()
	fmt.Printf("iV=%v Type=%T\n", iV, iV)
	switch iV.(type) {
	case Student:
		stu := iV.(Student)
		stu.Name = "Chloe Gan"
		fmt.Println(stu)
	}
}

func reflectNum(b interface{}) {
	rValue := reflect.ValueOf(b)
	rValue.Elem().SetInt(200)
}
