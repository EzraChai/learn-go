package main

import "fmt"

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
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000) //	Result
	exitChan := make(chan bool, 4)

	go putNumber(intChan, 8000)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("%v ", res)
	}
}

func putNumber(intChan chan int, n int) {
	for i := 0; i < n; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool

	for num := range intChan {
		flag = true
		if num == 1 || num == 0 {
			flag = false
		}
		for i := 2; i < num/2; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	//close(primeChan)
	exitChan <- true
}
