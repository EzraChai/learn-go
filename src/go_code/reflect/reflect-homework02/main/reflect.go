package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (cal Cal) GetSub(name string) {
	fmt.Printf("%v finish the subtraction %v - %v = %v", name, cal.Num1, cal.Num2, cal.Num1-cal.Num2)
}

func main() {
	cal := &Cal{
		Num1: 8,
		Num2: 3,
	}
	var name = "tom"

	reflectCal(*cal, name)
}

func reflectCal(b interface{}, name string) {
	rValue := reflect.ValueOf(b)

	numField := rValue.NumField()
	for i := 0; i < numField; i++ {
		fmt.Println(rValue.Type().Field(i).Name, ":", rValue.Field(i))
	}

	numMethod := rValue.NumMethod()
	var reflectArray []reflect.Value
	reflectArray = append(reflectArray, reflect.ValueOf(name))
	fmt.Println(reflectArray)
	for i := 0; i < numMethod; i++ {
		rValue.Method(i).Call(reflectArray)
	}
}
