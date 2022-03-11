package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var guess int

	rand.Seed(time.Now().Unix())
	value := rand.Intn(100)

	i := 1
	for {
		if i > 10 {
			fmt.Println("Idiot...")
			return
		}
		fmt.Println("Guess an integer [0 - 100]")
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Please try again")
			continue
		}
		if guess < 0 || guess > 100 {
			fmt.Println("Please enter an integer from [0 -100]")
			continue
		}
		if guess == value {
			switch i {
			case 1:
				fmt.Println("You are an genius!")
			case 2, 3:
				fmt.Println("You are clever, better than me!")
			case 4, 5, 6, 7, 8, 9:
				fmt.Println("Normal")
			case 10:
				fmt.Println("Finally you get it")
			}
			return
		} else {
			if guess > value {
				fmt.Println("Lower a bit")
				i++
			} else {
				fmt.Println("Higher a bit")
				i++
			}
		}
	}
}
