package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type C A

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	var a A
	var b B

	//a = b

	//	If you really want to migrate a different type of struct, make sure that the field and type are the same
	//	BUt why??!!
	a = A(b)

	fmt.Printf("a: %v b: %v\n", a, b)

	var c C
	//	Possible
	//	remember to force change the type
	c = C(a)
	fmt.Println(c)

	monster := Monster{"BABI", 500, "kill people"}

	//	Serialization
	monsterJson, err := json.Marshal(monster) //json.Marshal used REFLECT
	if err != nil {
		errors.New("Invalid struct")
	}
	fmt.Println(string(monsterJson))
	fmt.Println(monster)
}
