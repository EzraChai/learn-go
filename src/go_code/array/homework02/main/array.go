package main

import "fmt"

func main() {
	numArr := [...]int{1, 2, 3, 4, 5, 90, 109, 23}

	var largestNum int
	var index int = -1
	for i := 0; i < len(numArr); i++ {
		if i == 0 {
			largestNum = numArr[i]
			index = i
		} else {
			if numArr[i] > largestNum {
				largestNum = numArr[i]
				index = i
			}
		}
	}
	fmt.Println("Largest num :", largestNum, "index :", index)
}
