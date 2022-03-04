package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println("b =", b)

	fmt.Println("b occupies ", unsafe.Sizeof(b), "bytes")

	//	Cannot use 1
	//b := 1;

	var str string = "Hello"
	//str[0] = "a" // Cannot change the str value in Go
	fmt.Println(str)

	str3 := `str2 := "abc\nabc"
	fmt.Println(str2)`
	fmt.Println(str3)

	var str5 string
	for i := 0; i < 3; i++ {
		str5 += "Chloe Gan "
	}
	fmt.Println(str5)
}
