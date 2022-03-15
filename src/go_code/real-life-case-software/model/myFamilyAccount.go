package model

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type familyAccount struct {
	details string
	balance float64

	money float64
	note  string
	input uint8
}

func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		details: "",
		balance: 10000.0,
		money:   0.0,
		note:    "",
		input:   0,
	}
}

func (familyAccount *familyAccount) MainMenu() {

	//	Show main menu
	fmt.Println("\n-------------Household income and expenditure accounting software-------------")
	fmt.Println()
	fmt.Println("                    1 Income and Expenditure Details")
	fmt.Println("                    2 Register income")
	fmt.Println("                    3 Register expenses")
	fmt.Println("                    4 Exit software")

	for {
		fmt.Print("\nPlease choose [1-4]:")
		_, err := fmt.Scanln(&familyAccount.input)
		fmt.Println()
		if err != nil {
			return
		}

		switch familyAccount.input {
		case 1:
			familyAccount.showDetails()
		case 2:
			familyAccount.income()
		case 3:
			familyAccount.expenses()
		case 4:
			familyAccount.exit()
		}
		familyAccount.input = 0
	}
}

func (familyAccount *familyAccount) showDetails() {
	fmt.Println("------------------------Current income and expenditure details record------------------------")
	if familyAccount.details == "" {
		fmt.Printf("\nThere is no current income and expenditure.\n")
	} else {
		fmt.Printf("\nIncome & Expenditure\tAccount Balance\t\tAmount of Income & Expenditure\t\tStatement\n")
		fmt.Println(familyAccount.details)
	}
}

func (familyAccount *familyAccount) income() {
	fmt.Printf("Amount of current income: $")
	_, err := fmt.Scanln(&familyAccount.money)
	if familyAccount.money <= 0 {
		familyAccount.income()
	}
	if err != nil {
		return
	}
	familyAccount.balance += familyAccount.money
	fmt.Printf("Curent income statement:  ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		familyAccount.note = scanner.Text()
		familyAccount.note = strings.Trim(familyAccount.note, " ")
	}

	strMoney := fmt.Sprintf("%v", familyAccount.money)
	if familyAccount.money < 1000 {
		strMoney += fmt.Sprintf("\t")
	}

	familyAccount.details += fmt.Sprintf("Income\t\t\t$%v\t\t\t + $%v\t\t\t\t%v\n", familyAccount.balance, strMoney, familyAccount.note)
	familyAccount.clearInput()

}

func (familyAccount *familyAccount) expenses() {
	fmt.Printf("Amount of current expenses: $")
	_, err := fmt.Scanln(&familyAccount.money)
	if familyAccount.money <= 0 {
		familyAccount.expenses()
	}
	if err != nil {
		return
	}

	if familyAccount.money > familyAccount.balance {
		fmt.Println("Insufficient balance")
		return
	}

	//	Deduct the money from the balance
	familyAccount.balance -= familyAccount.money
	fmt.Printf("Curent expenses statement:  ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		familyAccount.note = scanner.Text()
		familyAccount.note = strings.Trim(familyAccount.note, " ")
	}

	strMoney := fmt.Sprintf("%v", familyAccount.money)
	if familyAccount.money < 1000 {
		strMoney += fmt.Sprintf("\t")
	}

	familyAccount.details += fmt.Sprintf("Expenses"+
		"\t\t$%v\t\t\t - $%v\t\t\t\t%v\n", familyAccount.balance, strMoney, familyAccount.note)
	familyAccount.clearInput()
}

func (familyAccount *familyAccount) exit() {
	userInput := "n"
	for {
		fmt.Printf("Are you sure you wanna quit this software? [ y/N ] : ")
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			return
		}
		userInput = strings.ToLower(userInput)
		if userInput == "y" || userInput == "n" {
			break
		}
		fmt.Println("Please try again...")
	}
	switch userInput {
	case "n":
		return
	case "y":
		os.Exit(0)
	default:
		fmt.Println("Please enter the correct option")
	}
}

func (familyAccount *familyAccount) clearInput() {
	familyAccount.note = ""
	familyAccount.money = 0.0
}
