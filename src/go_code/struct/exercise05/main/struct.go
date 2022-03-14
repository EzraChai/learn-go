package main

import "fmt"

type Visitor struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (visitor *Visitor) getTicketPrice() {
	if visitor.Age > 18 && visitor.Age <= 150 {
		fmt.Printf("%v's age is %v, ticket price = [RM%v]\n", visitor.Name, visitor.Age, 20)
	} else if visitor.Age >= 0 {
		fmt.Printf("%v's age is %v, ticket price = [RM%v]\n", visitor.Name, visitor.Age, 0)
	}
}

func main() {
	var v1 Visitor
	fmt.Println("Please enter your name:")
	fmt.Scanln(&v1.Name)
	fmt.Println("Please enter your age:")
	fmt.Scanln(&v1.Age)

	v1.getTicketPrice()

}
