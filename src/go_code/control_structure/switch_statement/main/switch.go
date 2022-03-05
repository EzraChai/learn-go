package main

import "fmt"

func test(key byte) byte {
	return key + 1
}

func main() {

	var key byte
	fmt.Println("Please enter a alphabet [a,b,c,d,e,f,g]")
	fmt.Scanf("%c", &key)

	switch test(key) {
	case 'a':
		fmt.Println("Monday")
	case 'b':
		fmt.Println("Tuesday")
	case 'c':
		fmt.Println("Wednesday")
	case 'd':
		fmt.Println("Thursday")
	case 'e':
		fmt.Println("Friday")
	case 'f':
		fmt.Println("Saturday")
	case 'g':
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a valid alphabet")
	}

	var n1 int32 = 20
	var n2 int32 = 20
	switch n1 {
	case n2, 4, 6:
		fmt.Println("OK")
	default:
		fmt.Println("NOT OK")
	}

	var age int = 10

	switch {
	case age == 10:
		fmt.Println("age == 10")
		fallthrough //weird lo
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("No such case")
	}

	var x interface{}
	var y int
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x types is %T", i)
	case int:
		fmt.Println("Type x is int")
	case float64:
		fmt.Println("Type x is float64")
	}
}
