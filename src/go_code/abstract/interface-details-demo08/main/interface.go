package main

import "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

type Student struct {
}

func (stu *Student) test01() {
	fmt.Println("Student test01()")
}
func (stu Student) test02() {
	fmt.Println("Student test02()")
}

//	Student does not implement AInterface (missing test03 method)
//func (stu Student) test03() {
//}

func (stu Student) test03() {
	fmt.Println("Student test03()")
}

//	interface{}
//	We can input every type of data into a function
type T interface {
}

type XInterface interface {
	test01()
}
type YInterface interface {
	test01()
}
type ZInterface interface {
	XInterface
	YInterface
}

func main() {
	var stu Student
	var a AInterface = &stu
	a.test03()
	a.test02()
	a.test01()

	var t T = stu
	fmt.Println(t)

	var num1 float64 = 6.43
	t = num1
	var z ZInterface = &stu
	z.test01()

	fmt.Println(t)
}
