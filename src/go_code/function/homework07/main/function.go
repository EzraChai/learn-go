package main

import (
	"errors"
	"fmt"
)

func main() {
	var inputNum, sum int
	for i := 1; i <= 5; i++ {
		fmt.Println("Please enter", i, "rational number")
		_, err := fmt.Scanln(&inputNum)
		if err != nil {
			fmt.Println("Please enter a rational number")
			i--
			continue
		}

		err2 := checkPrimeNumber(&inputNum)
		if err2 != nil {
			fmt.Println(err2)
			i--
			continue
		}

		sum += inputNum
	}
	fmt.Println("Sum:", sum)
}

func checkPrimeNumber(number *int) (err error) {
	isPrime := true
	if *number == 0 || *number == 1 {
	} else {
		for i := 2; i <= *number/2; i++ {
			if *number%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime == true {
			return nil
		}
	}
	return errors.New("invalid prime number")

}

//func addNumber(num ...int) {
//	sum := 0
//	for i := 0; i < len(num); i++ {
//		sum += num[i]
//	}
//	fmt.Println(sum)
//}
