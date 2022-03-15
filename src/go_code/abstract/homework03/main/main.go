package main

import (
	"fmt"
	"math/rand"
	"sort"
	"stpd.com/model"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var students model.Students
	for i := 1; i <= 30; i++ {
		student := model.Student{
			Name:  "student" + strconv.Itoa(i),
			Age:   17,
			Score: float64(rand.Intn(100)),
		}
		students = append(students, student)
	}
	fmt.Println(students)
	fmt.Printf(" NAME \t\t AGE \t SCORE \n")
	for _, student := range students {
		fmt.Printf("[%v] \t %v \t %v \n", student.Name, student.Age, student.Score)
	}

	sort.Sort(students)

	fmt.Printf("\n********************\n\n")

	fmt.Printf(" NAME \t\t AGE \t SCORE \n")
	for _, student := range students {
		fmt.Printf("[%v] \t %v \t %v \n", student.Name, student.Age, student.Score)
	}
}
