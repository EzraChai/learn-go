package main

import "fmt"

func main() {
	var arr [4][4]int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Scanln(&arr[i][j])
		}
	}

	printBox(arr)
	fmt.Println()

	for i := 0; i < len(arr)/2; i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j], arr[len(arr)-1-i][j] = arr[len(arr)-1-i][j], arr[i][j]
		}
	}
	printBox(arr)

}

func printBox(arr [4][4]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}
