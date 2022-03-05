package main

import "fmt"

func main() {
	var allClassSum float64
	var classNum int = 3
	var studentNum int = 5

	for i := 0; i < classNum; i++ {
		var sum float64

		fmt.Printf("-----------CLASS %v-----------\n\n", i+1)
		for i := 0; i < studentNum; i++ {
			var score float64
			fmt.Println("PLease enter", i+1, "student's mark:")
			fmt.Scanln(&score)
			sum += score
		}
		fmt.Printf("The  class %v's average marks: %v \n\n", i+1, sum/float64(studentNum))
		allClassSum += sum

	}
	fmt.Printf("Every class Sum is %v and average is %v", allClassSum, allClassSum/float64(studentNum*classNum))
}
