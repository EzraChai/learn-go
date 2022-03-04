package main

import (
	"fmt"
	"unsafe"
)

//int types
func main() {
	//	int int8 int16 int32 int64

	var i int = 1
	fmt.Println("i =", i)

	// constant 128 overflows int8
	// var j int8 = 128
	// fmt.Println(j)

	//	uint8 uint16 uint32 uint64

	//constant -12 overflows uint8
	// var j uint8 = -12
	// fmt.Println(j)

	//	int uint rune byte
	var num = 100
	fmt.Printf("The type of num is %T", num)

	var n2 int64 = 10
	fmt.Printf("The type of num2 is %T and num2 occupy %d bytes.\n", n2, unsafe.Sizeof(n2))

	var age uint8 = 21
	fmt.Printf("The type of num2 is %T and num2 occupy %d bytes.", age, unsafe.Sizeof(age))
}
