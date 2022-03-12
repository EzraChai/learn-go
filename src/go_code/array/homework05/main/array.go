package main

import "fmt"

func main() {
	var arr [3][5]int

	sum := 0.0
	for i, ints := range arr {
		classSum := 0
		for i2, i3 := range ints {
			fmt.Printf("Please enter %v class's %v student's mark\n", i+1, i2+1)
			fmt.Scanln(&i3)
			classSum += i3
		}
		fmt.Println("Class", i, "Average marks :", classSum/len(arr[i]))
		sum += float64(classSum)
	}
	fmt.Println("Average form 's marks :", sum/float64(len(arr)*len(arr[0])))
}
