package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(time.Now())
	currentTime := time.Now()
	fmt.Println(currentTime.Year())
	fmt.Println(currentTime.Month())
	fmt.Println(currentTime.Day())
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())

	timeList := strings.Split(currentTime.Format(time.UnixDate), " ")
	fmt.Println(timeList)

	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Println(i)
		time.Sleep(time.Millisecond * 10)
		i++
	}

	//	Usage of Unix and UnixNano
	fmt.Println("Unix:", time.Now().Unix(), "Date:", time.Now().Day())
}
