package main

import "fmt"

type LearnEnglishInterface interface {
	learningEnglish()
}

type Athlete struct {
	Name string
	Age  int
}

type BasketballPlayer struct {
	Athlete
}

type FootballPlayer struct {
	Athlete
}

type Student struct {
	Name string
	Age  string
}

type UniversityStudent struct {
	Student
}

type SecondarySchoolStudent struct {
	Student
}

func (footballPlayer *FootballPlayer) learningEnglish() {
	fmt.Println("Football Player: [" + footballPlayer.Name + "] is learning English")
}
func (universityStudent *UniversityStudent) learningEnglish() {
	fmt.Println("University Student: [" + universityStudent.Name + "] is learning English")
}

func main() {
	us := UniversityStudent{Student{
		Name: "Chloe Gan",
		Age:  "20",
	}}
	us.learningEnglish()
}
