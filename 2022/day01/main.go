package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calcElveCalories(input string) []int {
	elves := strings.Split(input, "\n\n")

	sum := make([]int, len(elves))

	for i, elve := range elves {
		calories := strings.Split(elve, "\n")

		for _, calorie := range calories {
			cal, _ := strconv.Atoi(calorie)
			sum[i] += cal
		}

		fmt.Printf("Elve #%d has %d calories\n", i, sum[i])
	}

	return sum
}

func findMax(elveCalories []int) int {
	max := elveCalories[0]

	for _, caloriesSum := range elveCalories {
		if caloriesSum > max {
			max = caloriesSum
		}
	}

	return max
}

func main() {
	elveCalories := calcElveCalories(input)
	max := findMax(elveCalories)
	fmt.Printf("Highest amount of calories: %d\n", max)
}
