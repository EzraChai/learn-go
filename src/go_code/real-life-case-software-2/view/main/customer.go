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

func (customerView *customerView) showEditCustomer() {
	var id int
	fmt.Print("Enter the customer's ID [-1 exit]: ")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	index := customerView.customerService.FindById(&id)
	if index == -1 {
		fmt.Println("Customer not found.")
		fmt.Println()
		customerView.showEditCustomer()
	}
	customer := customerView.customerService.GetCustomer(index)
	fmt.Print("Name [" + customer.Name + "]: ")
	scanner := bufio.NewScanner(os.Stdin)
	name := ""
	for scanner.Scan() {
		name = scanner.Text()
		break
	}
	if name != "" {
		customer.Name = name
	}
	for {
		fmt.Print("Gender [ Male/Female ] [" + customer.Gender + "]: ")
		gender := ""
		for scanner.Scan() {
			gender = scanner.Text()
			break
		}
		if gender != "" {
			oldGender := customer.Gender
			customer.Gender = gender
			if customer.ValidateGender() {
				break
			} else {
				customer.Gender = oldGender
			}
		} else {
			break
		}

	}

	for {
		fmt.Print("Age [" + strconv.Itoa(customer.Age) + "]: ")
		age := ""
		_, err := fmt.Scanln(&age)

		if err == nil || err.Error() == "unexpected newline" {
			if age == "" {
				break
			}
			newAge, err := strconv.Atoi(age)
			if err == nil {
				oldAge := customer.Age
				customer.Age = newAge
				if !customer.ValidateAge() {
					customer.Age = oldAge
					continue
				}
				break
			}
		}

	}
	fmt.Print("Phone [" + customer.Phone + "]: ")
	phone := ""
	for scanner.Scan() {
		phone = scanner.Text()
		break
	}
	if phone != "" {
		customer.Phone = phone
	}
	fmt.Print("Email [" + customer.Email + "]: ")
	email := ""
	for scanner.Scan() {
		email = scanner.Text()
		break
	}
	if email != "" {
		customer.Email = email
	}
	fmt.Println()
	customerView.customerService.EditCustomer(&customer, index)
}

func (customerView *customerView) showDeleteCustomer() {
	var id int

	fmt.Println("----------------Delete a customer----------------")
	fmt.Print("Enter the customer's ID [-1 exit]: ")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println()
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
			fmt.Println()
			fmt.Println("-----------------Delete complete-----------------")
			fmt.Println()
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
			customerView.showEditCustomer()
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
