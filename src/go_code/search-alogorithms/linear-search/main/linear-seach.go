package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var inputName string
	var nameArr = []string{"ezra chai", "chloe gan", "jack", "bob"}
	fmt.Println("Enter a name")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputName = scanner.Text()
	}
	lowerInputName := strings.ToLower(inputName)
	index := linearSearch(nameArr, lowerInputName)
	if index < 0 {
		fmt.Println("Name not found")
	} else {
		fmt.Println("Name found at index", index, "value :", nameArr[index])
	}
}

func linearSearch(nameArr []string, name string) (index int) {
	for i, s := range nameArr {
		if name == s {
			return i
		}
	}
	return -1
}
