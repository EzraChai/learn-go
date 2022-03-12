package main

import "fmt"

func main() {
	str := "juanzhe@helo"
	slice := str[:7]
	//	juanzhe
	fmt.Println(slice)

	//str[0] = 'z'
	//	"juanzhe@helo" => "cuanzhe@helo"
	//	This method can only change english word
	strSlice := []byte(str)
	strSlice[0] = 'c'
	str = string(strSlice)
	fmt.Println(str)

	strRune := []rune(str)
	strRune[0] = 'c'
	str = string(strRune)
	fmt.Println(str)

}
