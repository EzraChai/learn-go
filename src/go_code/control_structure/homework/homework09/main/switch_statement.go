package main

import (
	"fmt"
	"strings"
)

//	Waiting For Love
//	RIP AVICII
func main() {
	var currentDay string
	fmt.Println("What is the current Day?")
	fmt.Scanln(&currentDay)

	switch strings.ToLower(currentDay) {
	case "monday":
		fmt.Println("Left me broken")
	case "tuesday":
		fmt.Println("I was through with hoping")
	case "wednesday":
		fmt.Println("my empty arms were open")
	case "thursday":
		fmt.Println("waiting for love, waiting for love")
	case "friday":
		fmt.Println("Thank the stars")
	case "saturday":
		fmt.Println("I'm burning like a fire gone wild")
	case "sunday":
		fmt.Println("Guess I won't be comming to church")
	}

}
