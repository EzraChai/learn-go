package main

import "fmt"

//	array
func main() {
	heros := [...]string{"Batman", "Superman", "Ironman"}
	for _, hero := range heros {
		fmt.Printf("%v ", hero)
	}
}
