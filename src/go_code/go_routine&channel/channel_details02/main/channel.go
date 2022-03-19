package main

import "fmt"

func main() {
	//	 select can solve the problem, when the channel stuck while getting a data

	intChan := make(chan int, 10)
	for i := 1; i <= cap(intChan); i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 1; i <= cap(stringChan); i++ {
		stringChan <- "Hello " + fmt.Sprintf("%d", i)
	}

	getData(intChan, stringChan)
}

func getData(intChan chan int, stringChan chan string) {
	for {
		select {
		case v := <-intChan: //If the channel doesn't close, it will jam and causes deadlock
			fmt.Printf("Read data from intChan, data: %v \n", v)
		case v := <-stringChan:
			fmt.Printf("Read data from StringChan, data: %v \n", v)
		default:
			fmt.Println("Can't get any data")
			return
		}
	}
}
