package main

import "fmt"

func main() {
	n1 := 10
	n2 := -20

	swap(&n1, &n2)

	fmt.Println("n1 =", n1, "n2 =", n2)
}

//	EASY
func swap(n1, n2 *int) {
	*n1 += *n2
	*n2 = *n1 - *n2
	*n1 -= *n2
}
