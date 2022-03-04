package main

import "fmt"

//decimal
func main() {
	//oat32 float64

	var price float32 = 89.12
	fmt.Println("Price =", price)

	//while using float, it will cause loss of precision
	var num2 float64 = -7654630457.34235245
	fmt.Println("num2 =", num2)

	var num3 float32 = -123.0000091
	var num4 float64 = -123.0000091
	fmt.Println("num3 =", num3, "num4 =", num4)

	num5 := 1.1
	fmt.Printf("The type of num5 is %T", num5)

	num8 := 5.2342e31
	fmt.Println("num8 =", num8)
}
