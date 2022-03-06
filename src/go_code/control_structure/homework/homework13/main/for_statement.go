package main

import (
	"fmt"
)

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

	/**
	OUTPUT:

		*
	   * *
	  *   *
	 *     *
	*       *
	 *     *
	  *   *
	   * *
		*

	*/
	var count int = 5

	for a := 0; a < 2; a++ {
		if a == 0 {
			for i := 1; i <= count; i++ {
				for k := 1; k <= count-i; k++ {
					fmt.Print(" ")
				}
				for j := 1; j <= 2*i-1; j++ {
					if j == 1 || j == 2*i-1 {
						fmt.Print("*")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
		} else {
			for i := count - 1; i >= 0; i-- {
				for k := 1; k <= count-i; k++ {
					fmt.Print(" ")
				}
				for j := 1; j <= 2*i-1; j++ {
					if j == 1 || j == 2*i-1 {
						fmt.Print("*")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
		}
	}

}
