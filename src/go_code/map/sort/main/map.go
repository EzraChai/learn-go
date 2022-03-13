package main

import (
	"fmt"
)

func main() {
	map1 := make(map[int]int, 10)
	map1[10] = 10
	map1[0] = 90
	map1[1] = 40
	map1[2] = 43
	modify(map1)
	//	Auto extend
	map1[20] = 20
	fmt.Println(map1)

	//	NICE !!!
	type Student struct {
		Name    string
		Age     int
		Address string
	}

	StudentInfo := make(map[string]Student)

	stu1 := Student{Name: "Chloe Gan", Age: 16, Address: "Dataran Segar"}
	stu2 := Student{Name: "Ezra Chai", Age: 17, Address: "Port Dickson"}
	StudentInfo["1001"] = stu1
	StudentInfo["1002"] = stu2

	for id, student := range StudentInfo {
		fmt.Printf("ID [%v] name: %v age: %v address: %v\n", id, student.Name, student.Age, student.Address)
	}
}

func modify(map1 map[int]int) {
	map1[10] = 90
}
