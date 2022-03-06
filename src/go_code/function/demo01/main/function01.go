package main

import (
	"fmt"
	"stpd.com/utils"
)

func main() {

	var num1, num2, ans float64
	var operator byte
	fmt.Println("PLease enter a number")
	fmt.Scanln(&num1)
	fmt.Println("PLease enter a number")
	fmt.Scanln(&num2)
	fmt.Println("PLease enter an operator")
	fmt.Scanf("%c", &operator)

	ans = utils.Calculate(num1, num2, operator)

	fmt.Printf("%v %c %v = %v", num1, operator, num2, ans)

}
