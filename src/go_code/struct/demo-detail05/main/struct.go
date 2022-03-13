package main

import "fmt"

type Person struct {
	Name string `json:"name"`
}

//	like =>> this.xx
func (p Person) printName() {
	fmt.Println("printName() =>", p.Name)
}

func main() {
	var p1 Person
	p1.Name = "Chloe Gan"
	p1.printName()

}
