package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	score int
}

func (student *Student) ShowInfo() {
	fmt.Printf("Name:[%v] Age:[%v] Score:[%v]\n", student.Name, student.Age, student.GetScore())
}

func (student *Student) SetScore(score int) {
	if score < 0 || score > 100 {
		panic(errors.New("invalid score"))
	}
	student.score = score
}

func (student *Student) GetScore() int {
	return student.score
}

type Pupil struct {
	Student
}

type Graduate struct {
	Student
}

func main() {
	pupil := &Pupil{Student{
		Name: "Chloe Gan",
		Age:  16,
	}}
	fmt.Println(pupil.Student.Name)
	pupil.ShowInfo()
	pupil.SetScore(100)
	pupil.ShowInfo()
	fmt.Println(pupil.Student.Name)

	fmt.Println("***************")

	graduate := &Graduate{Student{
		Name: "Ezra Chai",
		Age:  20,
	}}

	graduate.ShowInfo()
	graduate.SetScore(99)
	graduate.ShowInfo()
}
