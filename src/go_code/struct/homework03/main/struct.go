package main

import "fmt"

type arrFunction struct {
}

func (*arrFunction) flipArray(arr *[3][3]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i])/2; j++ {
			arr[j][i], arr[i][j] = arr[i][j], arr[j][i]
		}
	}
}

func (*arrFunction) printArray(arr *[3][3]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}
func main() {
	var fun arrFunction
	arr := [3][3]int{}
	counter := 0
	for i := 0; i < len(arr); i++ {
		for k := 0; k < len(arr[i]); k++ {
			counter++
			arr[i][k] = counter
		}
	}
	fun.printArray(&arr)
	fun.flipArray(&arr)
	fmt.Println()
	fun.printArray(&arr)
}
