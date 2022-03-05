package main

import "fmt"

func main() {

	/*
			1. 	***
				***
				***

			2.	*
				**
				***

			3.	*
		 	   ***
		      *****
	*/
	var count int = 9

	for i := 1; i <= count; i++ {
		//for k := 2*(count-1) - i; k > 0; k -= 2 {
		//	fmt.Print(" ")
		//}
		for k := 1; k <= count-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == count {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()

	}
}
