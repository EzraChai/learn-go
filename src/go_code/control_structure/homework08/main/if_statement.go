package main

import "fmt"

func main() {
	var currentMonth int
	fmt.Println("Enter the current Month etc: 1,2,3,4,5")
	fmt.Scanln(&currentMonth)

	switch currentMonth {
	case 3, 4, 5:
		fmt.Println("Spring")
	case 6, 7, 8:
		fmt.Println("Summer")
	case 9, 10, 11:
		fmt.Println("Autumn")
	case 12, 1, 2:
		fmt.Println("Winter")
	default:
		fmt.Println("Invalid Month")
	}

}
