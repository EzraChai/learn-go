package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//	1.
	s := "Chloe Gan"
	fmt.Println(len(s)) //	9

	//	2.
	s2 := "Hello美女"
	//	If there is chinese word, change the type to rune to fix this issue => r := []rune(str)
	r := []rune(s2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	for _, i2 := range s2 {
		fmt.Printf("%c\n", i2)
	}

	//	3.
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Type: %T Value: %v\n", n, n)
	}

	//	4.
	str := strconv.Itoa(12345)
	fmt.Println(str)

	//	5.
	var bytes = []byte("Hello World")
	fmt.Printf("%v", bytes)
	fmt.Println()

	//	6.
	str2 := string([]byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100})
	fmt.Println(str2)

	str = strconv.FormatInt(123, 2)
	fmt.Println("The base 2 of number [123] is ", str)

	str = strconv.FormatInt(123, 16)
	fmt.Println("The base 16 of number [123] is ", str)

	//	7.
	b := strings.Contains("Chloe Gan", "Chloe")
	fmt.Println(b)

	//	8.
	num := strings.Count("cheese", "e")
	fmt.Println("Num of e in the word cheese is", num)

	//	9.
	b = strings.EqualFold("abc", "Abc")
	fmt.Println(b)

	//	10.
	fmt.Println(strings.TrimSpace("    Hello World    "))

	//	11.
	s = "Ezra Chai"
	numFront := strings.Index(s, "a")
	numBack := strings.LastIndex(s, "a")
	fmt.Println(numFront)
	fmt.Println(numBack)
	fmt.Printf("Character = %c\n", s[numFront])

	//	12.
	//	1.	String	2.	Old	Word	3.	New Word	4.	How many times you want to replace in a strings (n = -1 ==> Change all)
	fmt.Println(strings.Replace("Hello Chibai", "Chibai", "Chloe", -1))
}
