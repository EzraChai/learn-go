package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float64
	Sex   string
}

func (person Person) Print() {
	fmt.Println("------start------")
	fmt.Println(person)
	fmt.Println("-------end-------")
}

func (person Person) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (person Person) Set(name string, age int, score float64, sex string) {
	person.Name = name
	person.Age = age
	person.Score = score
	person.Sex = sex
}

func TestStruct(b interface{}, boolChan chan<- bool) {
	rType := reflect.TypeOf(b)
	rValue := reflect.ValueOf(b)
	rType = rType.Elem()
	rValue = rValue.Elem()
	kind := rValue.Kind()
	if kind != reflect.Struct {
		panic("expected struct")
	}
	fmt.Println("Reflect Type:", rType)
	fmt.Println("Reflect Value:", rValue)

	num := rValue.NumField()
	fmt.Printf("struct has %d fields\n", num)
	//	Iterate the structs' field
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: Value= %v\n", i, rValue.Field(i))
		if rType.Field(i).Name == "Name" {
			rValue.Field(i).SetString("Chloe Gan")
		}

		//	Get Struct tag, need to use reflect.Type to get the tag value
		tagValue := rType.Field(i).Tag.Get("json")
		if tagValue != "" {
			fmt.Printf("Field %d: tag=%v\n", i, tagValue)
		}
	}

	numOfMethod := rValue.NumMethod()
	fmt.Printf("struct has %d methods\n\n", numOfMethod)
	rValue.Method(1).Call(nil)

	//	Invoke the first method
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rValue.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())

	boolChan <- true
}

func main() {
	var boolChan chan bool
	boolChan = make(chan bool, 1)
	p1 := &Person{
		Name:  "chloegan",
		Age:   16,
		Score: 0,
	}
	go TestStruct(p1, boolChan)
	for {
		value := <-boolChan
		if value {
			break
		}
	}
	fmt.Println(p1)
}
