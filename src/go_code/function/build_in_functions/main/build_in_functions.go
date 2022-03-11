package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("Type: %T, value: %v, pointer: %v\n", num1, num1, &num1)
	num2 := new(int)
	fmt.Printf("Type: %T, value: %v, pointer: %v, pointer value: %v\n", num2, num2, &num2, *num2)
	*num2 = 200
	fmt.Println("value :", *num2)

}
