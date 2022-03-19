package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var num int
	numype := reflect.TypeOf(num)
	fmt.Println(numType.Name())

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("")
	}
}

func sayHello() {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error :", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang"
}
