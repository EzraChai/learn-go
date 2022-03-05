package main

import "fmt"

/**
SLVING PROBLEM TIPS
1）	先易后难，即将一个复杂的问题分解成简单的问题
2）	先死后活
*/
func main() {

	var allClassSum float64
	var classNum int = 3
	var studentNum int = 5
	var passCount int = 0

	for i := 0; i < classNum; i++ {
		var sum float64

		fmt.Printf("-----------CLASS %v-----------\n\n", i+1)
		for i := 0; i < studentNum; i++ {
			var score float64
			fmt.Println("PLease enter", i+1, "student's mark:")
			fmt.Scanln(&score)

			//	Checif the score is passed
			if score >= 60 {
				passCount++
			}
			sum += score
		}
		fmt.Printf("The  class %v's average marks: %v \n\n", i+1, sum/float64(studentNum))
		allClassSum += sum

	}
	fmt.Printf("Every class Sum is %v and average is %v\n", allClassSum, allClassSum/float64(studentNum*classNum))
	fmt.Println("Pass:", passCount)
}
