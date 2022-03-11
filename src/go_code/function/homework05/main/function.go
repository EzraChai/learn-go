package main

import (
	"errors"
	"fmt"
)

func printDay(year, month *int) {
	switch *month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 days")
	case 4, 6, 9, 11:
		fmt.Println("30 days")
	case 2:
		if *year%4 == 0 {
			fmt.Println("29 days")
		} else {
			fmt.Println("28 days")
		}
	}
	fmt.Println()
}

func checkValidMonth(month *int) (err error) {
	if *month > 12 || *month < 1 {
		return errors.New("invalid input")
	}
	return nil
}

func main() {
	var month, year int
	for {

		fmt.Println("Please enter the year")
		_, yearErr := fmt.Scanln(&year)

		if yearErr != nil {
			err := errors.New("invalid year")
			panic(err)
		}

		fmt.Println("Please enter the month")
		fmt.Scanln(&month)

		err := checkValidMonth(&month)
		if err != nil {
			panic(err)
		}
		printDay(&year, &month)

	}

}
