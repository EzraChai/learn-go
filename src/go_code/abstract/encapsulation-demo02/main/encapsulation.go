package main

import (
	"fmt"
	"person.com/model"
)

func main() {
	person := model.NewPerson("Chloe Gan")

	//	Encapsulation
	//person.Age = 18

	person.SetAge(18)
	person.SetSal(30000)

	fmt.Println(person.Name, " age:", person.GetAge(), "sal:", person.GetSalary())
}
