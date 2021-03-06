package main

import (
	"fmt"
	"strconv"
)

type Boy struct {
	No int
	// Name    string
	NextBoy *Boy
}

//	Returs the head of the list
func AddBoy(num int) *Boy {

	firstBoy := &Boy{}
	currentBoy := &Boy{}

	if num < 1 {
		return firstBoy
	}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		if i == 1 {
			firstBoy = boy
			currentBoy = boy
			currentBoy.NextBoy = firstBoy
		} else {
			currentBoy.NextBoy = boy
			currentBoy = boy
			currentBoy.NextBoy = firstBoy
		}
	}
	return firstBoy
}

func ShowBoyBoy(firstBoy *Boy) {

	if firstBoy.NextBoy == nil {
		fmt.Errorf("There's only one boy boy")
	}
	currentBoy := firstBoy

	for {

		fmt.Printf("Boy's Id: [" + strconv.Itoa(currentBoy.No) + "]\n")

		currentBoy = currentBoy.NextBoy

		if currentBoy == firstBoy {
			return
		}
	}

}

func main() {
	first := AddBoy(5)
	ShowBoyBoy(first)
}
