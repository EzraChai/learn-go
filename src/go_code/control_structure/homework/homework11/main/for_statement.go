package main

import "fmt"

func main() {

	var num int = 6

	for i := 0; i <= num; i++ {
		fmt.Println(i, "+", num-i, "=", num)
	}

}
