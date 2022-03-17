package main

import (
	"testing"
)

//	go test -v
//
//➜  testcase git:(main) ✗ go test -v
//=== RUN   TestAddUpper
//    cal_test.go:13: AddUpper(10) executing correctly
//--- PASS: TestAddUpper (0.00s)
//=== RUN   TestBinarySearch
//    cal_test.go:22: NIce
//--- PASS: TestBinarySearch (0.00s)
//PASS
//ok      stpd.com/testcase       0.003s

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10) Error, ExpectedValue = %v, CurrentValue = %v\n", 55, res)
		t.Fatalf("AddUpper(10) Error, ExpectedValue = %v, CurrentValue = %v", 55, res)
	}
	t.Logf("AddUpper(10) executing correctly")
}

//	go test -v -test.run TestBinarySearch
//	Test one unit Test
func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 213, 300}
	index := BinarySearch(arr, 4)
	if index != 3 {
		t.Fatalf("Expected value = 3, current value = %v", index)
	}
	t.Logf("NIce")
}
