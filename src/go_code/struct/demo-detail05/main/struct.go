package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string `json:"name"`
}

//	 like =>> this.xx	BUT NOT EXACTLY IS
func (person Person) printName() Person {
	fmt.Println("printName() =>", person.Name)
	person.Name = "Hi, " + person.Name
	return person
}

func main() {
	time.Now()
	var p1 Person
	p1.Name = "Chloe Gan"
	p1 = p1.printName()
	fmt.Println(p1.Name)

}
