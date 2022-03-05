package main

import (
	"fmt"
	"math"
)

//	The best function I ever made
func main() {
	var a, b, c float64
	a = 1
	b = 4
	c = -4

	m := (b * b) - 4*(a*c)
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		x2 := (-b - math.Sqrt(m)) / (2 * a)
		fmt.Println("Root is ", x1, "and ", x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		fmt.Println("Root is ", x1)
	} else {
		fmt.Println("No root")
	}
}
