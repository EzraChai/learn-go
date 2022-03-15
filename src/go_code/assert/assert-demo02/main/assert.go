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
func (p *Phone) Call() {
	fmt.Println("Phone is calling")
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

func (computer *Computer) working(usb USB) {
	usb.Start()
	//	Assert	//Useful
	phone, ok := usb.(*Phone)
	if ok {
		phone.Call()
	}
	usb.Stop()
}

func (computer *Computer) connecting(usb [3]USB) {
	for i := 0; i < len(usb); i++ {
		usb[i].Start()
		usb[i].Stop()
		fmt.Println()
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
	//comp.connecting(usbArr)
	for _, usb := range usbArr {
		comp.working(usb)
		fmt.Println()
	}
}
