package main

import "fmt"

type Usb interface {
	Start()
	Stop()
	//	Phone does not implement Usb (missing df method)
	//	df()
}

type Phone struct {
}

func (phone Phone) Start() {
	fmt.Println("Phone start working")
}

func (phone Phone) Stop() {
	fmt.Println("Phone stop working")
}

type Camera struct {
}

func (camera Camera) Start() {
	fmt.Println("Camera start working")
}

func (camera Camera) Stop() {
	fmt.Println("Camera stop working")
}

type Computer struct {
}

func (c *Computer) Working(usb Usb) {

	//	Using usb interface to run start and stop
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
}
