package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 3.142
	var b bool = true
	var myChar byte = 'h'

	var str string

	//Method 1 to transform the type to String
	str = fmt.Sprintf("%d", num1)
	fmt.Println(str)
	str2 := fmt.Sprintf("%f", num2)
	fmt.Println(str2)
	str3 := fmt.Sprintf("%t", b)
	fmt.Println(str3)
	str4 := fmt.Sprintf("%c", myChar)
	fmt.Println(str4)

	str5 := strconv.FormatInt(int64(num1), 10)
	fmt.Println(str5)
	// 'f' => FORMAT , 10 => 10 DECIMAL PLACES ,64 => THIS NUMBER IS A TYPE OF FLOAT64
	str6 := strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Println(str6)
	str7 := strconv.FormatBool(b)
	fmt.Println(str7)
	fmt.Println(strconv.Itoa(num1))
}
