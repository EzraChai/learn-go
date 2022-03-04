package main

import "fmt"

func main() {
	var days int = 97
	var day, week int

	week = days / 7
	day = days % 7
	fmt.Println(week, "weeks,", day, "day")
}
