package main

import "fmt"

func main() {
	var year int
	fmt.Println("Please enter a year")
	//	Get the years
	fmt.Scanln(&year)

	//	kinda weird because year 200 does not work so the logic is wrong I think
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		fmt.Println(year, "is a leap year")
	}
}
