package main

import "fmt"

type MethodUtils struct {
}

func (method *MethodUtils) printShape() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (*MethodUtils) area(width, height float64) float64 {
	return width * height
}

func (*MethodUtils) printShapeWithParams(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (*MethodUtils) oddOrEvenNumber(n int) {
	if n%2 == 0 {
		fmt.Println(n, "is a even number")
	} else {
		fmt.Println(n, "is a odd number")
	}
}

func (*MethodUtils) printShapeWithManyOption(row, column int, character byte) {
	for i := 0; i < column; i++ {
		for i := 0; i < row; i++ {
			fmt.Printf("%c", character)
		}
		fmt.Println()
	}
}

func (*MethodUtils) printMultiplicationTable(num int) {
	for j := 1; j <= num; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%v x %v = %v\t", i, j, j*i)
		}
		fmt.Println()
	}

}

func main() {
	var mu MethodUtils

	mu.printShape()

	fmt.Println()

	mu.printShapeWithParams(2, 6)

	area := mu.area(5, 2)
	fmt.Println("area :", area)

	mu.oddOrEvenNumber(10)

	mu.printShapeWithManyOption(5, 3, '@')

	mu.printMultiplicationTable(9)
}
