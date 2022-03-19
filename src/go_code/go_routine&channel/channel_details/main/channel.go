package main

import "fmt"

func main() {
	//var chan1 chan int	//	can read can write
	//var chan2 chan<- int //	can only write
	//var chan3 <-chan int

	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go receive(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
}

func receive(ch <-chan int, exitChan chan struct{}) {
	for value := range ch {
		fmt.Println(value)
	}
	var a struct{}
	exitChan <- a
}

func send(ch chan<- int, exitChan chan struct{}) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}
