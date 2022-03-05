package main

import "fmt"

func main() {
	var age int

	fmt.Println("What is your age?")
	_, err := fmt.Scanln(&age)
	if err != nil {
		return
	}

	if age > 18 {
		fmt.Println("You should take care of yourself.")
	} else {
		fmt.Println("Your parent should take care of you.")
	}

}
