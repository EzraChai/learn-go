package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v float64 = 1.2
	reflectFloat64(v)
}

func reflectFloat64(b interface{}) {
	rValue := reflect.ValueOf(b)
	fmt.Println(rValue.Type())
	fmt.Println(rValue.Kind())
	rInterface := rValue.Interface()

	value, ok := rInterface.(float64)
	if ok {
		fmt.Println(value)
	}
}
