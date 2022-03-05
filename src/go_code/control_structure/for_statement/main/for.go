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

	str := "我爱你 Chloe Gan"
	s := []rune(str)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println("")

	for index, value := range s {
		fmt.Printf("%v%c", index, value)
	}

	fmt.Println("")

	//	The is no while in GoLang

	x := 10
	for {
		if x == 0 {
			break
		}
		fmt.Println("Hello World")
		x--
	}

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
		fmt.Println()
	}
}
