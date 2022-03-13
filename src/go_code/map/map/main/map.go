package main

import "fmt"

//	map
func main() {
	var a map[string]string
	a = make(map[string]string, 10)

	fmt.Println(a)

	fmt.Println()

	cities := make(map[string]string)
	cities["no1"] = "Kuala Lumpur"
	cities["no2"] = "Seremban"
	fmt.Println(cities)

	fmt.Println()

	heroes := map[string]string{
		"hero1": "Superman",
		"hero2": "Batman",
	}
	heroes["hero3"] = "Ironman"
	fmt.Println(heroes)
}
