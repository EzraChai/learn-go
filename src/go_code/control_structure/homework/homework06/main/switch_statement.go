package main

import "fmt"

func main() {

	var alphabet byte

	fmt.Println("Please enter a alphabet [a,b,c,d,e]")
	fmt.Scanf("%c", &alphabet)

	switch alphabet {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("other")
	}
}
