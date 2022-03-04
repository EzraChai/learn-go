package main

import "fmt"

func main() {
	//2.5
	fmt.Println(10 / 3.0)

	//	2
	var n1 float32 = 10 / 4
	fmt.Println("n1:", n1)

	//	2.5
	var n2 float32 = 10.0 / 4
	fmt.Println("n2:", n2)

	//	Formula
	//	a % b = a - a / b * b
	fmt.Println("10 % 3 =", 10%3)     //	1
	fmt.Println("-10 % 3 =", -10%3)   //	-10 - (-10) / 3 * 3 = -1
	fmt.Println("10 % -3 =", 10%-3)   //	1
	fmt.Println("-10 % -3 =", -10%-3) //	-1

	var i int = 10
	i++
	fmt.Println(" i =", i)
	i--
	fmt.Println(" i =", i)
}
