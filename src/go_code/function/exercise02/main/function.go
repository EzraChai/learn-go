package main

import "fmt"

func main() {
	ans := f(5)
	fmt.Println(ans)
}

func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}
