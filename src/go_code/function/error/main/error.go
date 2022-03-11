package main

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		//err := recover() //	recover() is a built in function, can catch error
		if err := recover(); err != nil { //	There is an error

			//	This is where you handle error
			fmt.Println(err)
		}
	}()
	num1 := 100
	num2 := 0

	res := num1 / num2
	fmt.Println(res)
}

//	read the config file
//	if file name is invalid, throw an error
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	}
	return errors.New("invalid name for config file")
}

func test02() {
	err := readConf("config.in")
	if err != nil {
		//	if there is an error
		panic(err)
	}
	fmt.Println("The code continues")
}

func main() {
	test02()
	fmt.Println("Hello")
}
