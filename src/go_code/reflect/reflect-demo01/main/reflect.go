package main

import (
	"fmt"
	"reflect"
)

//	Use for reflect
func reflectTest01(b interface{}) {
	//Use reflect to get the variable's type and kind
	rType := reflect.TypeOf(b)
	fmt.Println("RType:", rType)

	rValue := reflect.ValueOf(b)

	//	rValue: 100 Type: reflect.Value
	//fmt.Printf("rValue: %v Type: %T", rValue, rValue)

	n2 := 2 * rValue.Int()
	fmt.Println(n2)

	//rValue.SetInt(n2)

	fmt.Printf("rValue: %v\n", rValue)

	//	Turn RValue to interface{}
	iV := rValue.Interface()

	num2 := iV.(int)
	fmt.Println("num2 =", num2)

}

//	REFLECT
func main() {
	//	declare variable num
	var num int = 100
	reflectTest01(num)
}
