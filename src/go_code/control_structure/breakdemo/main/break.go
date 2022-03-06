package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var times int
	rand.Seed(time.Now().UnixNano())
	var n int
label1:
	for {
		if n == 99 {
			break label1
		}
		n = rand.Intn(100) + 1
		times++
	}
	fmt.Println(times)
}
