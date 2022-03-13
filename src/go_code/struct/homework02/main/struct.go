package main

import "fmt"

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (calculator *Calculator) operator(operator byte) (ans float64) {
	switch operator {
	case '+':
		ans = calculator.Num1 + calculator.Num2
	case '-':
		ans = calculator.Num1 - calculator.Num2
	case '*', 'x':
		ans = calculator.Num1 * calculator.Num2
	case '/':
		ans = calculator.Num1 / calculator.Num2
	default:
		fmt.Println("Operator not found")
	}
	return
}

func main() {
	var cal Calculator
	cal.Num1 = 50
	cal.Num2 = 40
	ans := cal.operator('/')
	fmt.Println("Answer:", ans)
}
