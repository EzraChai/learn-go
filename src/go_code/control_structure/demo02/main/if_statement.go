package main

import "fmt"

func main() {
	var marks int
	fmt.Println("What is your GoLang Test Result?")
	fmt.Scanln(&marks)

	checkMarks(&marks)
}

func checkMarks(marks *int) {
	if *marks == 100 {
		fmt.Println("Present is a BMW car")
	} else if *marks >= 80 {
		fmt.Println("Present is a Iphone 7 plus")
	} else if *marks >= 60 {
		fmt.Println("Present is a iPad")
	} else {
		fmt.Println("No present")
	}
}
