package main

import (
	"account.com/model"
	"fmt"
)

func main() {
	account := model.NewAccount("chloegan", "050920", 25)
	fmt.Println(*account)
}
