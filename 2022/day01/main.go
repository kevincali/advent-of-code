package main

import (
	"fmt"
	"sort"
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

func main() {
	elveCalories := calcElveCalories(input)

	sort.Sort(sort.Reverse(sort.IntSlice(elveCalories)))

	fmt.Printf("Highest amount of calories: %d\n", elveCalories[0])

	topThree := elveCalories[0] + elveCalories[1] + elveCalories[2]
	fmt.Printf("Sum of the top three elves: %d\n", topThree)
}
