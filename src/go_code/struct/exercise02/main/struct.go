package main

import (
	"fmt"
	"math"
	"strconv"
)

type Circle struct {
	Radius int `json:"radius"`
}

func (circle Circle) area() float64 {
	return math.Pi * float64(circle.Radius*circle.Radius)
}

//	Calculate area of a circle
func main() {
	var c1 Circle
	c1.Radius = 4
	//	Find the area of the circle with the radius of 4 cm
	area := c1.area()

	//	Format the area to four decimal places
	areaInString := fmt.Sprintf("%.4f", area)
	//	Convert the string to float64
	area, _ = strconv.ParseFloat(areaInString, 64)
	//	Area = 50.2655cm^2, Radius = 4cm
	fmt.Println(area)
}
