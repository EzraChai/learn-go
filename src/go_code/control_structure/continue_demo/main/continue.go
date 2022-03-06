package main

import "fmt"

func main() {
	//for i := 0; i < 4; i++ {
	//	for j := 0; j < 10; j++ {
	//		if j == 2 {
	//			continue
	//		}
	//		fmt.Println("j=", j)
	//	}
	//}
	//for i := 0; i < 5; i++ {
	//	if i == 3 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	//
	//for h := 1; h <= 100; h++ {
	//	if h%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(h)
	//}
	//
	//var positiveNum int = 0
	//var negativeNum int = 0
	//var num int
	//for {
	//	fmt.Println("Please enter a integer")
	//	fmt.Scanln(&num)
	//	if num == 0 {
	//		break
	//	} else if num > 0 {
	//		positiveNum++
	//	} else {
	//		negativeNum++
	//	}
	//}
	money := 100000.0
	var count int
	for {
		if money < 1000 {
			break
		}
		if money > 50000 {
			money *= 0.95
		} else {
			money -= 1000.0
		}
		count++
	}

	fmt.Println("count:", count)
}
