package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		v := <-exitChan
		if v {
			break
		}
	}
	fmt.Println("Hello World!")
}

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for num := range intChan {
		fmt.Println(num)
	}
	exitChan <- true
	close(exitChan)
}
