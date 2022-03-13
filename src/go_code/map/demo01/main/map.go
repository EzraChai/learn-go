package main

import "fmt"

//	map
func main() {
	//	var (variable name) map[keyType]valueType
	var a map[string]string

	//	panic: assignment to entry in nil map
	//a[0] = "Chloe Gan"

	a = make(map[string]string, 10)
	a["love"] = "No one"
	a["love2"] = ""
	a["love"] = "Chloe Gan"

	fmt.Println(a)
}
