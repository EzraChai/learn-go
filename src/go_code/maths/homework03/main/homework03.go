package main

import "fmt"

func main() {

	a := 28
	b := 28
	c := 39

	fmt.Println(getLargerNumber(a, b))
	fmt.Println(getLargerNumberWithin3Number(a, b, c))

}

func getLargerNumber(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getLargerNumberWithin3Number(a int, b int, c int) int {
	if a > b {
		if a > c {
			return a
		}
	} else {
		if b > c {
			return b
		}
	}
	return c
}
