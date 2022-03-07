package main

import "fmt"

//	weird question lo
func main() {
	numPeach := consumePeach(1, 10)
	fmt.Println(numPeach)
}

func consumePeach(peach float64, days int) float64 {
	if days == 1 {
		return peach
	}
	peach = 2 * (peach + 1)
	return consumePeach(peach, days-1)
}
