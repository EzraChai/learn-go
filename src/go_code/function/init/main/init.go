package main

import (
	"fmt"
	"stpd.com/utils"
)

var PI = getPi()

func getPi() (PI float64) {
	PI = 3.142
	fmt.Println("GLOBAL VARIABLE INIT")
	return
}

func init() {
	fmt.Println("Invoke init()")
}

// Declare global variables =>> init() =>> main()
func main() {
	fmt.Println("GLOBAL VARIABLE NAME =", utils.Name)
	fmt.Println("Invoke Main()")
}
