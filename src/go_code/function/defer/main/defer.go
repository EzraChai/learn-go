package main

import (
	"fmt"
)

func sum(n1, n2 int) int {
	//	When application is executing at the line defer, it will create another independent stack (defer)
	//	When this function is finish, it will run according first in, last out
	defer fmt.Println("ok1 n1 =", n1) //	defer
	defer fmt.Println("ok2 n2 =", n2) //	defer
	//	defer file.Close()
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res =", res)
	return res
}

func main() {
	/*
		OUTPUT:
			ok3 res = 32
			ok2 n2 = 20
			ok1 n1 = 10
			32
	*/
	res := sum(10, 20)
	fmt.Println(res)
}
