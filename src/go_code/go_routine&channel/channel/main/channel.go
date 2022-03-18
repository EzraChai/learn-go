package main

import "fmt"

func main() {
	//	DEMO the usagen of Channel

	var intChan chan int
	intChan = make(chan int, 3)

	//	0xc000102000 =>> address
	fmt.Println(intChan)

	intChan <- 10
	intChan <- 211
	intChan <- 200
	//	fatal error: all goroutines are asleep - deadlock!
	//	intChan <- 400

	fmt.Println("Length:", len(intChan), "Capacity:", cap(intChan)) //	3,	3

	n2 := <-intChan
	n3 := <-intChan
	n4 := <-intChan

	fmt.Println(n2, n3, n4)

	//n5 := <-intChan

	//	fatal error: all goroutines are asleep - deadlock!
	//fmt.Println(n2, n3, n4, n5)

	intChan <- 23
	//<-intChan
	n5 := <-intChan
	fmt.Println(n5)

	fmt.Println("Length:", len(intChan), "Capacity:", cap(intChan)) //	2,	3

}
