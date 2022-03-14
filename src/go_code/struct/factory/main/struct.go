package main

import (
	"fmt"
	"stpd.com/model"
)

func main() {
	//	Return a pointer of the student struct
	student1 := model.NewStudent("Chloe Gan", 100)
	fmt.Println(*student1)
	fmt.Println("Name:", student1.Name)
	fmt.Println("Score:", student1.GetScore())
	student1.SetScore(99)
	fmt.Println("Score:", student1.GetScore())

}
