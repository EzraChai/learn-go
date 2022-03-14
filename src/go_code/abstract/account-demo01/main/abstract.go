package main

import "fmt"

type Account struct {
	AccountNo string
	Password  string
	Balance   float64
}

func (account *Account) deposit(money float64, pwd string) {
	if account.Password == pwd {
		if money > 0 {
			account.Balance += money
		}
	}
}

//	Withdraw money from the account
func (account *Account) withdraw(money float64, pwd string) {
	if account.Password == pwd {
		if account.Balance >= money {
			account.Balance -= money
		}
	}
}

func (account *Account) query(pwd string) float64 {
	if account.Password == pwd {
		return account.Balance
	}
	return -1
}

func main() {
	var account = &Account{
		AccountNo: "1001",
		Password:  "chloegan",
		Balance:   100,
	}
	fmt.Println(account.query("chloegan"))
	account.deposit(10000, "chloegan")
	account.withdraw(900, "chloegan")
	fmt.Println(account.query("chloegan"))

}
