package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool

	//	strconv.ParseBool(str) function returns two value which is (value bool, err error)
	//	I want to get only value bool and I dont want to get the err, so use _ to ignore
	b, _ = strconv.ParseBool(str)
	fmt.Printf("%v %T\n", b, b)

	var str2 string = "123456"
	num, _ := strconv.ParseInt(str2, 10, 32)
	fmt.Printf("%v %T\n", num, num)

	var str3 string = "123.456"
	num2, _ := strconv.ParseFloat(str3, 64)
	num3 := float32(num2)
	fmt.Printf("%v %T\n", num3, num3)

	var str4 string = "hello"
	n3, err := strconv.ParseInt(str4, 10, 32)
	fmt.Println(err)
	fmt.Printf("%v %T\n", n3, n3)
}
