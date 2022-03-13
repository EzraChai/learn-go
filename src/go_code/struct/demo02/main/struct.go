package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //Pointer
	slice  []int             //Slice
	map1   map[string]string //Map
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	//	1. method
	var p1 Person

	//	{ 0 [0 0 0 0 0] <nil> [] map[]}
	fmt.Println("P1 =", p1)
	if p1.ptr == nil {
		fmt.Println("OK")
	}
	if p1.slice == nil {
		fmt.Println("OK")
	}
	if p1.map1 == nil {
		fmt.Println("OK")
	}

	//	2. method
	p2 := Person{Name: "Chloe Gan", Age: 16}
	fmt.Println("P2 =", p2)

	//	3. method
	p3 := new(Person)
	//	(*p3).Name = "Ezra Chai" == p3.Name = "Ezra Chai"
	p3.Name = "Ezra Chai"
	fmt.Println("P3 =", *p3)

	//	4. method
	var p4 *Person = &Person{Name: "Jack"}
	p4.Age = 2
	fmt.Println("P4 =", p4)

	//	Must use make() to use slice
	p1.slice = make([]int, 2)
	p1.slice[0] = 2
	fmt.Println(p1.slice)

	p1.map1 = make(map[string]string)
	p1.map1["1001"] = "Chloe Gan"
	fmt.Println(p1.map1)

	var monster1 Monster

	monster1.Name = "BAbI"
	monster1.Age = 500

	monster2 := monster1
	monster2.Name = "CHIBAI"

	fmt.Println(monster1) //	{BAbI 500}
	fmt.Println(monster2) //	{CHIBAI 500}
}
