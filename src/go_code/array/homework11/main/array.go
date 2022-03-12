package main

import "fmt"

func main() {
	arr := [10]string{"AA", "BB", "CC", "DD", "EE", "AA", "FF", "GG", "HH", "II"}
	b, index := findAA(&arr)
	fmt.Println("Is there any AA?", b, " Where it is?", index)
}

func findAA(arr *[10]string) (b bool, index []int) {
	for i := 0; i < len(arr); i++ {
		if "AA" == arr[i] {
			index = append(index, i)
			if !b {
				b = true
			}
		}
	}
	return
}
