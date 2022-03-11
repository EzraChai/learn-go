package main

import "fmt"

func main() {
	var alphabets [26]byte
	for i := 0; i < 26; i++ {
		alphabets[i] = byte('A' + i)
	}

	for i := 0; i < len(alphabets); i++ {
		fmt.Printf("%c ", alphabets[i])
	}
}
