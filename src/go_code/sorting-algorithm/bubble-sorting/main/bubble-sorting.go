package main

import "fmt"

func main() {
	arr := []int{24, 69, 80, 57, 13}

	bubbleSort(arr)

	fmt.Println(arr)
}

//	Sort from small number to big number
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			//	Check if the front number is bigger than the current Number
			//	If the front number is bigger than the current number, switch the position
			if arr[j-1] > arr[j] {
				//	Switch the position of the number
				arr[j-1] += arr[j]
				arr[j] = arr[j-1] - arr[j]
				arr[j-1] -= arr[j]
			}
		}
	}
}
