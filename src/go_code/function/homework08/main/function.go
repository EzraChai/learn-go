package main

import "fmt"

func main() {
	rationalNumber()
	fmt.Println()
	atozZtoA()
}

func atozZtoA() {
	//	a == 97,z == 122,Z == 90 ,A == 64
	for i := 97; i <= 122; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println()
	for i := 90; i > 64; i-- {
		fmt.Printf("%c ", i)
	}
}

func rationalNumber() {
	counter := 0
	for i := 2; i <= 100; i++ {
		primeNum := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				primeNum = false
				break
			}
		}
		if primeNum {
			counter++
			if counter%5 == 0 {
				fmt.Printf("%v\t\n", i)
			} else {
				fmt.Printf("%v\t", i)
			}
		}
	}
}
