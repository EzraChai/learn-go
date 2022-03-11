package main

import "fmt"

//	array
func main() {
	var intArr [3]int
	fmt.Println(&intArr)

	//	0xc00001a240, intArr[0] 0xc00001a240,intArr[1] 0xc00001a248
	fmt.Printf("%p, intArr[0] %p,intArr[1] %p", &intArr, &intArr[0], &intArr[1])
}
