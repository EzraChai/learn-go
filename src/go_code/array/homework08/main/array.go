package main

import (
	"errors"
	"fmt"
)

func main() {
	var arr [3][4]int

	//	Get the input from the user
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			_, e := fmt.Scanln(&arr[i][j])
			if e != nil {
				panic(errors.New("please enter only number"))
			}
		}
	}

	//	Clear the number around the outer box
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == 0 || i == len(arr)-1 || j == 0 || j == len(arr[i])-1 {
				arr[i][j] = 0
			}
		}
	}

	//	Print the array out
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}
