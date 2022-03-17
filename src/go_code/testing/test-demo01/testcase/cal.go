package main

func addUpper(n int) (res int) {
	for i := 1; i <= n; i++ {
		res += i
	}
	return
}

func BinarySearch(arr []int, n int) int {
	leftIndex := 0
	rightIndex := len(arr) - 1
	for {
		if leftIndex > rightIndex {
			return -1
		}
		middleIndex := (leftIndex + rightIndex) / 2
		if arr[middleIndex] > n {
			rightIndex = middleIndex - 1
		} else if arr[middleIndex] < n {
			leftIndex = middleIndex + 1
		} else {
			return middleIndex
		}
	}
}

func getSum(n1, n2 int) int {
	return n1 + n2
}
