package main

import "fmt"

func main() {
	//	1. method
	var intArr [4]int
	var slice01 = intArr[:]
	fmt.Println(slice01)

	//	2. method
	var slice02 = make([]float64, 2)
	slice02[1] = 100
	fmt.Println(slice02)

	//	3. method
	slice03 := []string{"Chloe", "Gan"}
	fmt.Println(slice03)

	//	for
	for i := 0; i < len(slice03); i++ {
		fmt.Println(slice03[i])
	}
	fmt.Println()
	//	for - range
	for _, s := range slice03 {
		fmt.Println(s)
	}

	var slice3 []int = []int{100, 200, 300}
	slice3 = append(slice3, 400, 500)

	var appendSlice []int = []int{600, 700}
	slice3 = append(slice3, appendSlice...)
	fmt.Println(slice3)

	//var slice4 = make([]int, len(slice3))
	var slice4 = make([]int, 10)
	copy(slice4, slice3)
	fmt.Println(slice4)
}
