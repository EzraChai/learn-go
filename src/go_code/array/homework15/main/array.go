package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var arr [8]int

	insertRandomNumberIntoArray(&arr)
	fmt.Println("before:", arr)

	newArr := removeHighestNumAndLowestNum(&arr)
	fmt.Println("after:", newArr)

	avgScore := getAvgScore(newArr)
	fmt.Println("Avg Score:", avgScore)

	bestJudge, worstJudge := getBestJudgeAndWorstJudge(newArr, avgScore)
	fmt.Println("Best Judge =", bestJudge, " Worst Judge =", worstJudge)

}

func insertRandomNumberIntoArray(arr *[8]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10)
	}
}

func removeHighestNumAndLowestNum(arr *[8]int) (newArr []int) {
	highestNumIndex := 0
	lowestNumIndex := 0

	for i := 0; i < len(arr); i++ {
		if arr[highestNumIndex] < arr[i] {
			highestNumIndex = i
		} else if arr[lowestNumIndex] > arr[i] {
			lowestNumIndex = i
		}
	}
	fmt.Println("Judge ", lowestNumIndex, "and", highestNumIndex)

	for i := 0; i < len(arr); i++ {
		if i != highestNumIndex && i != lowestNumIndex {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

func getAvgScore(arr []int) (avgScore float64) {
	sum := 0
	totalJudge := len(arr)
	for i := 0; i < totalJudge; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(totalJudge)
}

func getBestJudgeAndWorstJudge(arr []int, avgScore float64) (bestJudge, worstJudge []int) {
	//lowestDifferenceJudge := -1
	var arrDiference []float64
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arrDiference = append(arrDiference, -1)
			continue
		}
		if float64(arr[i]) < avgScore {
			arrDiference = append(arrDiference, avgScore-float64(arr[i]))
		} else {
			arrDiference = append(arrDiference, float64(arr[i])-avgScore)
		}
	}

	smallestIndex := 0
	biggestIndex := 0
	for i := 0; i < len(arrDiference); i++ {
		if arrDiference[smallestIndex] > arrDiference[i] && arrDiference[i] >= 0 {
			smallestIndex = i
		} else if arrDiference[biggestIndex] < arrDiference[i] && arrDiference[i] >= 0 {
			biggestIndex = i
		}
	}
	for i := 0; i < len(arrDiference); i++ {
		if arrDiference[smallestIndex] == arrDiference[i] {
			bestJudge = append(bestJudge, i)
		} else if arrDiference[biggestIndex] == arrDiference[i] {
			worstJudge = append(worstJudge, i)
		}
	}

	return
}
