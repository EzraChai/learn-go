package main

import (
	"fmt"
	"strings"
)

//	Cool and useful function
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	fun := makeSuffix(".jpg")
	s1 := fun("123")
	s2 := fun("chloe_gan.jpg")
	fmt.Println(s1)
	fmt.Println(s2)
}
