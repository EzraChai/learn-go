package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Salary   float64 `json:"salary"`
	Skill    string  `json:"skill"`
}

func testStruct() []byte {
	p1 := Person{
		Name:     "Chloe Gan",
		Age:      16,
		Birthday: "20-09-2005",
		Salary:   30000,
		Skill:    "Read books",
	}
	marshal, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	return marshal
}

func testMap() []byte {
	map1 := make(map[string]interface{}, 2)
	map1["person1"] = Person{
		Name:     "Chloe Gan",
		Age:      16,
		Birthday: "20-09-2005",
		Salary:   30000,
		Skill:    "Read books",
	}
	map1["person2"] = Person{
		Name:     "Ezra Chai",
		Age:      17,
		Birthday: "25-02-2005",
		Salary:   32000,
		Skill:    "Coding",
	}

	marshal, err := json.Marshal(map1)
	if err != nil {
		panic(err)
	}
	return marshal
}

func testInt() []byte {
	var num1 float64 = 3.142
	marshal, err := json.Marshal(num1)
	if err != nil {
		panic(err)
	}
	return marshal
}

func testSlice() []byte {
	var slice []map[string]interface{}
	aMap := make(map[string]interface{}, 2)
	bMap := make(map[string]interface{}, 2)
	aMap["name"] = "jack"
	aMap["age"] = 16
	bMap["name"] = "tan"

	slice = append(slice, aMap)
	slice = append(slice, bMap)

	//	Serialize the slice
	marshal, err := json.Marshal(slice)
	if err != nil {
		panic(err)
	}
	return marshal
}

//	serialize to JSON format
func main() {
	jsonPerson := testStruct()
	fmt.Println(string(jsonPerson))

	jsonMap := testMap()
	fmt.Println(string(jsonMap))

	jsonSlice := testSlice()
	fmt.Println(string(jsonSlice))

	jsonInt := testInt()
	fmt.Println(string(jsonInt))
}
