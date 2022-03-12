package main

import "fmt"

func main() {
	slice := fbn(20)
	fmt.Println(slice)
}

func fbn(n int) (fbnSlice []uint64) {
	fbnSlice = make([]uint64, n)
	fbnSlice[0] = 1

	if n == 1 {
		return
	}
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
