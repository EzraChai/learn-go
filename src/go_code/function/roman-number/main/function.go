package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func convertHundred(numberStrArr string) (str string) {
	switch numberStrArr {
	case "9":
		str += "CM"
	case "8", "7", "6", "5":
		str += mapAllNum678(numberStrArr, "D", "C")
	case "4":
		str += "CD"
	case "3", "2", "1":
		str += mapAllNum123(numberStrArr, "C")
	}
	return str
}

func convertTen(numberStrArr string) (str string) {
	switch numberStrArr {
	case "9":
		str += "XC"
	case "8", "7", "6", "5":
		str += mapAllNum678(numberStrArr, "L", "X")
	case "4":
		str += "XL"
	case "3", "2", "1":
		str += mapAllNum123(numberStrArr, "X")
	}
	return str
}

func convertOne(numberStrArr string) (str string) {
	switch numberStrArr {
	case "9":
		str += "IX"
	case "8", "7", "6", "5":
		str += mapAllNum678(numberStrArr, "V", "I")
	case "4":
		str += "IV"
	case "3", "2", "1":
		str += mapAllNum123(numberStrArr, "I")
	}
	return str
}

func mapAllNum123(numberStrArr, letter string) (str string) {
	num, _ := strconv.Atoi(numberStrArr)
	for {
		if num == 0 {
			break
		}
		str += letter

		num -= 1
	}
	return str
}

func mapAllNum678(numberStrArr, frontLetter, letter string) (str string) {
	num, _ := strconv.Atoi(numberStrArr)
	str += frontLetter
	for {
		if num == 5 {
			break
		}
		str += letter
		num--
	}
	return str
}

func convertIntToRomanNumber(number int) (str string) {
	//	Check if the number is a positive number
	if number < 0 {
		panic(errors.New("only positive number between 1 to 1499 is available"))
	}

	//	Turn number to string
	numberStr := strconv.Itoa(number)

	//	Turn number to array [n1, n2, n3, n4]
	numberStrArr := strings.Split(numberStr, "")

	if len(numberStrArr) == 4 { //	Example [2,2,2,2]
		num, _ := strconv.Atoi(numberStrArr[0])
		for {
			if num == 0 {
				break
			}
			str += "M"
			num -= 1
		}
		str += convertHundred(numberStrArr[1])
		str += convertTen(numberStrArr[2])
		str += convertOne(numberStrArr[3])
	} else if len(numberStrArr) == 3 { //	Example [2,2,2]
		str += convertHundred(numberStrArr[0])
		str += convertTen(numberStrArr[1])
		str += convertOne(numberStrArr[2])
	} else if len(numberStrArr) == 2 { //	Example [2,2]
		str += convertTen(numberStrArr[0])
		str += convertOne(numberStrArr[1])
	} else { //	Example [2]
		str += convertOne(numberStrArr[0])
	}
	return
}

//	Function to convert an Integer between 1 and 1499 to a Roman Number
func main() {
	num := 1499
	str := convertIntToRomanNumber(num)
	fmt.Println(str)
}
