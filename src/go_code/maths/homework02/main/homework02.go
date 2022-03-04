package main

import (
	"fmt"
	"stpd.com/temperature"
)

func main() {

	a := 9
	b := 2

	a += b
	b = a - b
	a -= b
	fmt.Println(a, b)

	celcius := temperature.FahrenheitToCelcius(90)
	fmt.Println(celcius, "F")
}
