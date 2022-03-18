package main

import (
	"fmt"
)

func main() {
	numChan := make(chan int, 2000)
	resChan := make(chan map[string]int, 2000)
	closeChan := make(chan bool, 8)
	doneCount := 0
	number := 1
	totalGoRoutineCount := 8
	var array = make([]string, 2000)

	go addNumber(numChan, 2000)

	for i := 0; i < totalGoRoutineCount; i++ {
		go calculate(numChan, resChan, closeChan)
	}
	for {
		v := <-closeChan
		if v {
			doneCount++
			if doneCount == totalGoRoutineCount {
				close(closeChan)
				close(resChan)
				break
			}
		}
	}
	for value := range resChan {
		array[value["key"]-1] = fmt.Sprintf("res[%v] = %v", value["key"], value["value"])
		number++
	}
	for _, s := range array {
		fmt.Println(s)
	}
}

func addNumber(numChan chan int, n int) {
	for i := 1; i <= n; i++ {
		numChan <- i
	}
	close(numChan)
}

func calculate(numChan chan int, resChan chan map[string]int, closeChan chan bool) {
	for num := range numChan {
		mapNum := make(map[string]int, 2)
		sum := 0
		for i := 0; i <= num; i++ {
			sum += i
		}
		mapNum["key"] = num
		mapNum["value"] = sum
		resChan <- mapNum
	}
	closeChan <- true
}
