package main

import "fmt"

func main() {
	var marks [5]float64

	for i := 0; i < len(marks); i++ {
		fmt.Println("Enter", i+1, "student's mark:")
		fmt.Scanln(&marks[i])
	}
	for i := 0; i < len(marks); i++ {
		fmt.Printf("Student %v 's mark: %v\n", i+1, marks[i])
	}
}
