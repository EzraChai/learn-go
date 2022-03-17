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

func deserializeStruct() {
	str := "{\"name\":\"Chloe Gan\",\"age\":16,\"birthday\":\"20-09-2005\",\"salary\":30000,\"skill\":\"Read books\"}"
	person1 := Person{}
	err := json.Unmarshal([]byte(str), &person1)
	if err != nil {
		panic(err)
	}
	fmt.Println(person1)
}

func deserializeMap() {
	str := "{\"person1\":{\"name\":\"Chloe Gan\",\"age\":16,\"birthday\":\"20-09-2005\",\"salary\":30000,\"skill\":\"Read books\"},\"person2\":{\"name\":\"Ezra Chai\",\"age\":17,\"birthday\":\"25-02-2005\",\"salary\":32000,\"skill\":\"Coding\"}}"
	var map1 map[string]Person
	err := json.Unmarshal([]byte(str), &map1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", map1["person1"].Name)
}

func deserializeSlice() {
	str := "[{\"age\":16,\"name\":\"jack\"},{\"name\":\"tan\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		panic(err)
	}
	for _, m := range slice {
		fmt.Println(m)
		name := m["name"]
		fmt.Println(name)
	}
}

//	deserialize JSON format
func main() {
	deserializeStruct()
	deserializeMap()
	deserializeSlice()
}
