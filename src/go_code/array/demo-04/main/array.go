package main

import "fmt"

func test01(arr *[3]int) {
	arr[0] = 88
}

//	array
func main() {
	var arr01 [3]int
	arr01[0] = 1
	//	Different type (float64 != int64)
	//arr01[1] = 1.1
	//	Out of bound
	//arr01[3] = 1

	var arr02 [3]float32
	var arr03 [3]string
	var arr04 [3]bool
	fmt.Printf("%v %v %v\n", arr02, arr03, arr04)

	arr := [3]int{11, 22, 33}
	test01(&arr)
	fmt.Println(arr)
}
