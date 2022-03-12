package main

import "fmt"

func main() {
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	slice := intArr[1:3]
	fmt.Println(slice)
	fmt.Println("len =", len(slice))

	//	Initial Capacity = len(slice) * 2
	fmt.Println("cap =", cap(slice))

	//	slice[0]'s pointer: 0xc0000181b8, arr[1] pointer: 0xc0000181b8
	fmt.Printf("slice[0]'s pointer: %p, arr[1] pointer: %p\n", &slice[0], &intArr[1])
	slice[0] = 34
	fmt.Println(intArr[1])
}
