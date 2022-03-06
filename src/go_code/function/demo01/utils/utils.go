package utils

import "fmt"

func Calculate(num1 float64, num2 float64, operator byte) float64 {
	var ans float64
	switch operator {
	case '+':
		ans = num1 + num2
	case '-':
		ans = num1 - num2
	case '*':
		ans = num1 * num2
	case '/':
		ans = num1 / num2
	default:
		fmt.Println("False operator")
		return 0
	}
	return ans
}
