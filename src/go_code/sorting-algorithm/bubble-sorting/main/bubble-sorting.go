package main

import "fmt"

func main() {
	arr := []int{24, 69, 80, 57, 13}

	arr = bubbleSort(arr)
	fmt.Println(arr)
}

//	Sort from small number to big number
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		//	First number as larger number
		largerNum := arr[0]
		for j := 1; j < len(arr)-i; j++ {
			//	Check if the front number is bigger than the current Number
			//	If the front number is bigger than the current number, switch the position
			if largerNum > arr[j] {
				//	Switch the position of the number
				arr[j-1] += arr[j]
				arr[j] = arr[j-1] - arr[j]
				arr[j-1] -= arr[j]
			}
			//	change the larger number to the current Number
			largerNum = arr[j]
		}
	}
	return arr
}
