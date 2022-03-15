package main

import "fmt"

type USB interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p *Phone) Start() {
	fmt.Println("Phone is starting")
}
func (p *Phone) Stop() {
	fmt.Println("Phone is stoping")
}

type Camera struct {
}

func (c *Camera) Start() {
	fmt.Println("Camera is starting")
}
func (c *Camera) Stop() {
	fmt.Println("Camera is stoping")
}

type Computer struct {
}

func (computer *Computer) connecting(usb [3]USB) {
	for i := 0; i < len(usb); i++ {
		usb[i].Start()
		usb[i].Stop()
	}
}

func main() {
	// Polymorphism
	var usbArr [3]USB
	usbArr[0] = &Phone{}
	usbArr[1] = &Phone{}
	usbArr[2] = &Camera{}
	fmt.Println(usbArr)
	var comp Computer
	comp.connecting(usbArr)
}
