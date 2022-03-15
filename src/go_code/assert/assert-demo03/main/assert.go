package main

import "fmt"

type Student struct {
}

//	function that can
func TypeJudge(items ...interface{}) {
	for _, item := range items {
		switch item.(type) {
		case bool:
			fmt.Printf("Bool Type, Value = %v\n", item)
		case float32:
			fmt.Printf("float32 Type, Value = %v\n", item)
		case float64:
			fmt.Printf("float64 Type, Value = %v \n", item)
		case int, int64:
			fmt.Printf("int Type, Value = %v \n", item)
		case int32:
			fmt.Printf("int32 Type, Value = %v \n", item)
		case string:
			fmt.Printf("string Type, Value = %v \n", item)
		case Student:
			fmt.Printf("Student Type, Value = %v \n", item)
		case *Student:
			fmt.Printf("*Student Type, Value = %v \n", item)
		default:
			fmt.Printf("No such type, Value = %v \n", item)
		}
	}
}

func main() {
	var n1 float32 = 3.3
	var n2 float64 = 2.3
	n3 := 23
	home := "Malaysia"
	s1 := Student{}
	s2 := &Student{}

	TypeJudge(n1, n2, n3, home, s1, s2)
}
