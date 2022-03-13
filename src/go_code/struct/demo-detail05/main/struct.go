package main

import "fmt"

type Person struct {
	Name string `json:"name"`
}

//	 like =>> this.xx
func (p Person) printName() Person {
	fmt.Println("printName() =>", p.Name)
	p.Name = "Hi, " + p.Name
	return p
}

func main() {
	var p1 Person
	p1.Name = "Chloe Gan"
	p1 = p1.printName()
	fmt.Println(p1.Name)

}
