package main

import (
	"fmt"
	"unsafe"
)

type Node struct {
	row    int
	column int
	value  int
}

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1 //	black
	chessMap[2][3] = 2 //	white
	printTwoDimensionalArray(chessMap)

	var savedArray []Node

	//	Standard sparse array have recorded the two-dimensional array's structure( row column initial value)
	savedArray = append(savedArray, Node{
		row:    10,
		column: 10,
		value:  0,
	})

	for i := 0; i < len(chessMap); i++ {
		for f := 0; f < len(chessMap[f]); f++ {
			if chessMap[i][f] != 0 {
				savedArray = append(savedArray, Node{
					row:    i,
					column: f,
					value:  chessMap[i][f],
				})
			}
		}
	}
	fmt.Println()

	//	Iterate the saved Array
	fmt.Println("row\tcol\tvalue")
	for i := 0; i < len(savedArray); i++ {
		fmt.Printf("%v\t%v\t%v\n", savedArray[i].row, savedArray[i].column, savedArray[i].value)
	}

	var chessMap2 [11][11]int
	for i := 0; i < len(savedArray); i++ {
		chessMap2[savedArray[i].row][savedArray[i].column] = savedArray[i].value
	}
	fmt.Println()

	printTwoDimensionalArray(chessMap2)

	fmt.Println()

	//	Size of saved chess map before using sparse array
	fmt.Println("Size before compress ", unsafe.Sizeof(chessMap))

	//	Size of saved chess map after using sparse array
	fmt.Println("Size after compress ", unsafe.Sizeof(savedArray))
}

func printTwoDimensionalArray(array [11][11]int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("%v  ", array[i][j])
		}
		fmt.Println()
	}
}
