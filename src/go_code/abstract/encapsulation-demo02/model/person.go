package model

import "errors"

type person struct {
	Name   string
	age    int //other package cannot directly query it
	salary float64
}

//	New a person with the name
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//	Set a person's age
func (person *person) SetAge(age int) {
	if age >= 18 && age <= 150 {
		person.age = age
	} else {
		panic(errors.New("invalid age"))
	}
}

//	Get a person's age
func (person *person) GetAge() int {
	return person.age
}

//	Set a person's monthly salary
func (person *person) SetSal(salary float64) {
	if salary > 0 {
		person.salary = salary
	} else {
		panic(errors.New("invalid salary"))
	}
}

//	Get a person's salary
func (person *person) GetSalary() float64 {
	return person.salary
}
