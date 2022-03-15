package main

import (
	"errors"
	"fmt"
)

type Pupil struct {
	Name  string
	Age   int
	score int
}

//	Graduate, HighSchoolStudent ...

func (pupil *Pupil) ShowInfo() {
	fmt.Printf("Name:[%v] Age:[%v] Score:[%v]\n", pupil.Name, pupil.Age, pupil.GetScore())
}

func (pupil *Pupil) SetScore(score int) {
	if score < 0 || score > 100 {
		panic(errors.New("invalid score"))
	}
	pupil.score = score
}

func (pupil *Pupil) examination() {
	fmt.Println("Pupil is exam")
}

func (pupil *Pupil) GetScore() int {
	return pupil.score
}

func main() {
	pupil := &Pupil{
		Name: "Chloe Gan",
		Age:  16,
	}
	pupil.ShowInfo()
	pupil.SetScore(100)
	pupil.ShowInfo()
}
