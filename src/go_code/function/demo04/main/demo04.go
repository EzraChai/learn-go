package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFunc(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//	weird syntax but kinda nice
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//... is useful !
func sum(args ...int) (sum int) {
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

type myFunType func(int, int) int

func main() {

	a := getSum

	fmt.Printf("Type of a is %T, Type of getSum is %T\n", a, getSum)

	fmt.Println(myFunc(a, 40, 20))

	res := a(10, 20) // == res := getSum(10,20)
	fmt.Println("res:", res)

	type myInt int

	var num1 myInt
	num1 = 40
	fmt.Println("num1 =", num1)
	fmt.Printf("Type = %T\n", num1)

	fmt.Println(getSumAndSub(10, 5))

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}
