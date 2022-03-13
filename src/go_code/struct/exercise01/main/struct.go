package main

import "fmt"

type Person struct {
	Name string
}

func (person Person) speak() {
	fmt.Println(person.Name, "is a good person")
}

func (person Person) calculate() (sum int) {
	//	O(n)
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	return
}

func (person Person) calculate2(n int) (sum int) {
	//	O(n)
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

func (person Person) getSum(n1, n2 int) int {
	//	O(1)
	return n1 + n2
}

func main() {
	p1 := Person{"Jack"}
	p1.speak()

	sum := p1.calculate()
	fmt.Println("sum :", sum)

	sum2 := p1.calculate2(10)
	fmt.Println("sum2 :", sum2)

	fmt.Println(p1.getSum(10, 20))

}
