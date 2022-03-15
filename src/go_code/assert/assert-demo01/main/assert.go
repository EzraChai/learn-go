package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	var a interface{}

	var point = Point{
		x: 1,
		y: 2,
	}
	a = point

	var b Point
	//	cannot use a (type interface {}) as type Point in assignment: need type assertion
	b = a.(Point)
	fmt.Println(b)

	var num1 float32 = 3.4
	a = num1

	num2, flag := a.(float64)
	if flag {
		fmt.Println("Convert successful")
		fmt.Println(num2)
	} else {
		fmt.Println("Convert Fail")
	}
	fmt.Println("**********")
}
