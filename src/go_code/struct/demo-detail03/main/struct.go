package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rectangle struct {
	leftUp    Point
	rightDown Point
}

func main() {
	r1 := Rectangle{Point{1, 2}, Point{3, 4}}

	/*
		r1.leftUp.x POINTER=0xc000016140 + 8
		r1.leftUp.y POINTER=0xc000016148 + 8
		r1.rightDown.x POINTER=0xc000016150 + 8
		r1.rightDown.y POINTER=0xc000016158 + 8

		* base 16

		There are connection between all this initialize field which is Add 8 byte for the pointer
	*/
	fmt.Printf("r1.leftUp.x POINTER=%p \n", &r1.leftUp.x)
	fmt.Printf("r1.leftUp.y POINTER=%p \n", &r1.leftUp.y)
	fmt.Printf("r1.rightDown.x POINTER=%p \n", &r1.rightDown.x)
	fmt.Printf("r1.rightDown.y POINTER=%p \n", &r1.rightDown.y)

}
