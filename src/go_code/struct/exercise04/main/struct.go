package main

import "fmt"

type Box struct {
	Length float64 `json:"length"`
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

func (box *Box) volume() float64 {
	return box.Length * box.Height * box.Width
}

func main() {
	var box Box

	fmt.Println("Length")
	fmt.Scanln(&box.Length)
	fmt.Println("Height")
	fmt.Scanln(&box.Height)
	fmt.Println("Width")
	fmt.Scanln(&box.Width)

	volume := fmt.Sprintf("%.2f", box.volume())
	fmt.Println(volume, "cm^3")
}
