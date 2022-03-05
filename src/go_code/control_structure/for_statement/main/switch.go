package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		//fmt.Println("Hello, Chloe Gan")
	}

	k := 3
	for {
		if k == 0 {
			break
		}
		fmt.Println("I love Chloe Gan")
		k--
	}

	s := "Chloe Gan"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println("\n")

	/**
	Nice. I did it!!!
	Output:
	          *
	          ***
	         *****
	        *******
	       *********
	      ***********
	     *************
	    ***************
	   *****************
	  *******************

	*/
	for i := 1; i < 20; i += 2 {
		for j := 20; 2*j > i; j-- {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
