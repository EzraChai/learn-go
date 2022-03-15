package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//	Get the user's input
	var input uint8

	var detail string = ""
	var balance float64 = 10000.0

	//	Show main menu
	fmt.Println("\n-------------Household income and expenditure accounting software-------------")
	fmt.Println()
	fmt.Println("                    1 Income and Expenditure Details")
	fmt.Println("                    2 Register income")
	fmt.Println("                    3 Register expenses")
	fmt.Println("                    4 Exit software")

	for {
		var money float64 = 0.0
		var note string = ""

		fmt.Print("\nPlease choose [1-4]:")
		fmt.Scanln(&input)

		switch input {
		case 1:
			fmt.Println("------------------------Current income and expenditure details record------------------------")
			if detail == "" {
				fmt.Printf("\nThere is no current income and expenditure.\n")
			} else {
				fmt.Printf("\nIncome & Expenditure\tAccount Balance\t\tAmount of Income & Expenditure\t\tStatement\n")
				fmt.Println(detail)
			}
		case 2:
			fmt.Printf("Amount of current income: $")
			_, err := fmt.Scanln(&money)
			if err != nil {
				break
			}
			balance += money
			fmt.Printf("Curent income statement:  ")
			scanner := bufio.NewScanner(os.Stdin)

			if scanner.Scan() {
				note = scanner.Text()
				note = strings.Trim(note, " ")
			}

			strMoney := fmt.Sprintf("%v", money)
			if money < 1000 {
				strMoney += fmt.Sprintf("\t")
			}

			detail += fmt.Sprintf("Income\t\t\t$%v\t\t\t + $%v\t\t\t\t%v\n", balance, strMoney, note)
		case 3:
			fmt.Printf("Amount of current expenses: $")
			_, err := fmt.Scanln(&money)
			if err != nil {
				fmt.Println("Please enter number")
				break
			}

			if money > balance {
				fmt.Println("Insufficient balance")
				break
			}

			//	Deduct the money from the balance
			balance -= money
			fmt.Printf("Curent expenses statement:  ")
			scanner := bufio.NewScanner(os.Stdin)

			if scanner.Scan() {
				note = scanner.Text()
				note = strings.Trim(note, " ")
			}

			strMoney := fmt.Sprintf("%v", money)
			if money < 1000 {
				strMoney += fmt.Sprintf("\t")
			}

			detail += fmt.Sprintf("Expenses"+
				"\t\t$%v\t\t\t - $%v\t\t\t\t%v\n", balance, strMoney, note)
		case 4:
			userInput := "n"
			for {
				fmt.Printf("Are you sure you wanna quit this software? [ y/N ] : ")
				fmt.Scanln(&userInput)
				userInput = strings.ToLower(userInput)
				if userInput == "y" || userInput == "n" {
					break
				}
				fmt.Println("Please try again...")
			}
			switch userInput {
			case "n":
				break
			case "y":
				return
			}
		default:
			fmt.Println("Please enter the correct option")
		}
		input = 0
	}
}
