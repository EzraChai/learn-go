package main

import "fmt"

func main() {
	i := 100
	var n1 float32 = float32(i)
	fmt.Println(n1)
	fmt.Printf("i = %T, n1 = %T\n", i, n1)

	i2 := 129
	var i3 int8 = int8(i2)
	fmt.Println(i3)

}
