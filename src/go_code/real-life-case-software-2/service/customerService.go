package service

import (
	"customer.com/model"
	"fmt"
)

//	Customer Function including CRUD
type CustomerService struct {
	customers []model.Customer

	//	This field can become a new Customer's ID + 1
	customerNumber int
}

func (cs *CustomerService) AddCustomer(customer *model.Customer) {
	cs.customerNumber++
	customer.Id = cs.customerNumber
	cs.customers = append(cs.customers, *customer)
}

func (cs *CustomerService) ListCustomer() []model.Customer {
	return cs.customers
}

func (cs *CustomerService) FindById(id *int) (index int) {
	if *id > cs.customerNumber {
		return -1
	}
	leftIndex := 0
	rightIndex := cs.customerNumber
	for {
		if leftIndex > rightIndex {
			return -1
		}
		middleIndex := (leftIndex + rightIndex) / 2
		if cs.customers[middleIndex].Id < *id {
			leftIndex = middleIndex + 1
		} else if cs.customers[middleIndex].Id > *id {
			rightIndex = middleIndex - 1
		} else {
			return index
		}
	}
}

func (cs *CustomerService) DeleteByIndex(id *int) bool {
	index := cs.FindById(id)

	if index == -1 {
		fmt.Println("Customer not found.")
		fmt.Println()
		return false
	}
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

func NewCustomerService() *CustomerService {
	c := model.NewCustomer(1, "Chloe Gan", "Female", 16, "012-3456789", "chloegan.cg@gmail.com")

	customerService := &CustomerService{}
	customerService.customerNumber = 1
	customerService.customers = append(customerService.customers, *c)
	return customerService
}
