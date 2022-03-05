package main

import "fmt"

func main() {

	var name string
	var age int8
	var salary float32
	var passedExam bool

	//	fmt.Scanln
	//fmt.Println("Name")
	//fmt.Scanln(&name)
	//
	//fmt.Println("Age")
	//fmt.Scanln(&age)
	//
	//fmt.Println("Salary")
	//fmt.Scanln(&salary)
	//
	//fmt.Println("Passed Exam ?")
	//fmt.Scanln(&passedExam)
	//

	fmt.Println("Please enter your name, age, salary, passed exam or not. Please use space to separate it.")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &passedExam)

	fmt.Printf("Name\t\tAge\tSalary\t\tPassed Exam\n%q\t%v\t%f\t%v", name, age, salary, passedExam)

}
