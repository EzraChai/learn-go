package hello

import (
	"fmt"
	"strconv"
)

func HelloGo(i int) {
	fmt.Println("Hello Golang", strconv.Itoa(i))
	//time.Sleep(time.Second)
}

func HelloWorld(i int) {
	fmt.Println("Hello World", strconv.Itoa(i))
	//time.Sleep(time.Second)
}
