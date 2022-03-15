package main

import (
	"bufio"
	"customer.com/model"
	"customer.com/service"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type customerView struct {
	input           uint8
	toggleExit      string
	customerService *service.CustomerService
}

func (customerView *customerView) showAddCustomer() {
	customer := model.Customer{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("------------------Add a Customer------------------")
	fmt.Print("Name: ")
	for scanner.Scan() {
		customer.Name = scanner.Text()
		break
	}
	for {
		fmt.Print("Gender [ Male/Female ]: ")
		for scanner.Scan() {
			customer.Gender = scanner.Text()
			break
		}
		if customer.ValidateGender() {
			break
		}
	}

	for {
		fmt.Print("Age : ")
		_, err := fmt.Scanln(&customer.Age)
		if err == nil {
			if !customer.ValidateAge() {
				continue
			}
			break
		}
	}
	fmt.Print("Phone: ")
	for scanner.Scan() {
		customer.Phone = scanner.Text()
		break
	}
	fmt.Print("Email: ")
	for scanner.Scan() {
		customer.Email = scanner.Text()
		break
	}
	customerView.customerService.AddCustomer(&customer)

	fmt.Println()
}

func (customerView *customerView) showDeleteCustomer() {
	var id int

	fmt.Println("----------------Delete a customer----------------")
	fmt.Print("Enter the customer's ID [-1 exit]: ")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	customerView.toggleExit = ""

	for {
		fmt.Print("Are you sure you want to delete customer with id: [" + strconv.Itoa(id) + "] ? [y/N]")
		fmt.Scanln(&customerView.toggleExit)
		switch strings.Trim(strings.ToLower(customerView.toggleExit), " ") {
		case "y":
			if !customerView.customerService.DeleteByIndex(&id) {
				customerView.showDeleteCustomer()
			}
			fmt.Println("----------------Delete complete----------------")
			return
		case "n":
			return
		default:
			break
		}
	}
}

func (customerView *customerView) showListCustomer() {
	customers := customerView.customerService.ListCustomer()
	fmt.Println("-----------------Customers' List-----------------")
	fmt.Printf("\nID\tName\t\tGender\tAge\tPhone\t\tEmail\n")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println()
}

func (customerView *customerView) exit() {
	customerView.toggleExit = ""

	fmt.Println("-------------------Exit Program-------------------")
	fmt.Print("Are you really want to exit this program? [ y/N ] ")
	fmt.Scanln(&customerView.toggleExit)
	switch strings.Trim(strings.ToLower(customerView.toggleExit), " ") {
	case "y":
		os.Exit(0)
	case "n":
		fmt.Println()
		break
	default:
		customerView.exit()
	}
}

//	Show Main menu
func (customerView *customerView) showMenu() {
	for {
		fmt.Println("---------Customer information management software---------")
		fmt.Println("                 1 Add a Customer")
		fmt.Println("                 2 Edit a Customer")
		fmt.Println("                 3 Delete a Customer")
		fmt.Println("                 4 Customer's List")
		fmt.Println("                 5 Exit")
		fmt.Println()
		for {
			fmt.Print("Please enter [1-5]: ")
			_, err := fmt.Scanln(&customerView.input)
			if err == nil {
				break
			}
		}
		switch customerView.input {
		case 1:
			customerView.showAddCustomer()
		case 2:
			fmt.Println("                 2 Edit a Customer")
		case 3:
			customerView.showDeleteCustomer()
		case 4:
			customerView.showListCustomer()
		case 5:
			customerView.exit()
		}
	}
}

func main() {
	var cv customerView
	cv.customerService = service.NewCustomerService()
	cv.showMenu()
}
