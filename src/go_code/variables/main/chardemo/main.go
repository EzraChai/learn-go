package main

import "fmt"

//char
func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	// 'a' => 97
	// '0' => 48
	fmt.Println("c1 =", c1)
	fmt.Println("c2 =", c2)

	fmt.Printf("c1 = %c, c2 = %c\n", c1, c2)

	var c3 int = '晨'
	var c4 int = '纾'

	fmt.Printf("%c%c\n", c3, c4)

	var c5 int = 050225
	fmt.Printf("%c\n", c5)

	var n1 = 10 + 'a'
	fmt.Printf("%c", n1)
}
