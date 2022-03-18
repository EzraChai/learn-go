package main

import (
	"fmt"
	"runtime"
)

type Number struct {
	n int
}

//GOROUTINE
//	Characteristic of Coroutine
//	1.	Have independent Stack ace
//	2. 	Shared program heap space
//	3.	Controlled by the user
//	4.	Coroutine is lightweight thread

func main() {
	//n := Number{n: 10}
	//n.RationalNumber()

	num := runtime.NumCPU()
	fmt.Println(num)

}

func (number *Number) RationalNumber() {

}
