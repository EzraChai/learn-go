package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Age    int     `json:"age"`
	Score  float64 `json:"score"`
}

func (student *Student) say() string {
	return "ID: " + strconv.Itoa(student.Id) + " NAME: " + student.Name + " GENDER: " + student.Gender +
		" AGE: " + strconv.Itoa(student.Age) + " SCORE: " + fmt.Sprintf("%.2f", student.Score) + "%"
}

func main() {
	var s1 = Student{
		Id:     1,
		Name:   "Chloe Gan",
		Gender: "Woman",
		Age:    16,
		Score:  100,
	}
	fmt.Println(s1.say())

}
