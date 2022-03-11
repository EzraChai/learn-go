package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	str := ""
	currentUnixNano := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	afterUnixNano := time.Now().Unix()

	usedTime := afterUnixNano - currentUnixNano
	fmt.Println(usedTime)
}
