package model

import "errors"

type account struct {
	accountId string
	balance   float64
	password  string
}

func NewAccount(accountId string, password string, balance float64) *account {

	if len(accountId) < 6 || len(accountId) > 10 {
		panic(errors.New("length of account id must be between 6 and 10"))
	}
	if len(password) != 6 {
		panic(errors.New("password must be in 6 digit"))
	}

	if balance <= 20 {
		panic(errors.New("must deposit more than RM 20"))
	}

	return &account{
		accountId: accountId,
		password:  password,
		balance:   balance,
	}
}

func (acc *account) ChangePassword(password string) {
	if len(password) != 6 {
		panic(errors.New("password must be in 6 digit"))
	} else {
		acc.password = password
	}
}

func (acc *account) Deposit(money float64) {
	if money < 1 {
		panic(errors.New("invalid money"))
	} else {
		if acc.balance-money < 0 {
			panic(errors.New("insufficient money"))
		} else {
			acc.balance -= money
		}
	}
}
