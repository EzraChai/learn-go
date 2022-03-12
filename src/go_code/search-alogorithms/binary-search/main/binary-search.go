package main

import (
	"fmt"
)

func main() {
	var findNum = 1000
	var arr = []int{1, 8, 10, 89, 1000, 1234}

	index := binarySearchV2(arr, findNum)
	binarySearchV3(arr, 0, len(arr)-1, findNum)
	fmt.Println(index)
}

func binarySearch(arr []int, num int) (index int) {
	leftIndex := 0
	rightIndex := len(arr) - 1
	for {
		if rightIndex > leftIndex {
			return -1
		}
		middle := (leftIndex + rightIndex) / 2
		if arr[middle] < num {
			leftIndex = middle - 1
		} else if arr[middle] > num {
			rightIndex = middle + 1
		} else if arr[middle] == num {
			return middle
		}
	}
}

func binarySearchV2(arr []int, num int) (index int) {
	leftIndex := 0
	rightIndex := len(arr)
	middleIndex := 0
	for {
		if leftIndex > rightIndex {
			return -1
		}

		middleIndex = (leftIndex + rightIndex) / 2

		if arr[middleIndex] > num {
			rightIndex = middleIndex + 1
		} else if arr[middleIndex] < num {
			leftIndex = middleIndex - 1
		} else {
			return middleIndex
		}
	}
}

func binarySearchV3(arr []int, leftIndex int, rightIndex, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("not found")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if arr[middle] < findVal {
		binarySearchV3(arr, middle+1, rightIndex, findVal)
	} else if arr[middle] > findVal {
		binarySearchV3(arr, leftIndex, middle-1, findVal)
	} else {
		fmt.Println("index :", middle)
	}
}
