package model

import (
	"fmt"
	"strings"
)

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustomer(id int, name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (customer *Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", customer.Id, customer.Name, customer.Gender, customer.Age, customer.Phone, customer.Email)
}

func (customer *Customer) ValidateAge() bool {
	if customer.Age > 0 && customer.Age < 150 {
		return true
	}
	return false
}

func (customer *Customer) ValidateGender() bool {
	customer.Gender = strings.Trim(strings.ToLower(customer.Gender), " ")
	switch customer.Gender {
	case "male":
		customer.Gender = "Male"
		return true
	case "female":
		customer.Gender = "Female"
		return true
	default:
		return false
	}
}
