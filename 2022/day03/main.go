package main

import (
	"fmt"
	"regexp"
	"strings"
)

// a-z have priorities 1-26
// A-Z have priorities 27-52
// therefore offset the ASCII value to reach the corresponding prio value
func getPrio(letter string) int {
	ascii := int([]rune(letter)[0])

	lowercaseMatch, _ := regexp.MatchString("[a-z]", letter)
	if lowercaseMatch {
		return ascii - 96
	}

	uppercaseMatch, _ := regexp.MatchString("[A-Z]", letter)
	if uppercaseMatch {
		return ascii - 38
	}

	return -1
}

func prioritizeDuplicates(data string) int {
	rucksacks := strings.Fields(data)
	prioSum := 0

	for _, rucksack := range rucksacks {
		fmt.Println("\nChecking rucksack:", rucksack)

		middle := len(rucksack) / 2
		firstHalf := rucksack[:middle]
		secondHalf := rucksack[middle:]

		for _, rune := range firstHalf {
			letter := string(rune)

			if strings.ContainsAny(secondHalf, letter) {
				fmt.Println("Found duplicate item in both compartments:", letter)
				prioSum += getPrio(letter)

				break
			}
		}
	}

	fmt.Printf("\n\nPriority sum of duplicates is %d\n\n", prioSum)
	return prioSum
}

func prioritizeBadges(data string) int {
	rucksacks := strings.Fields(data)
	prioSum := 0

	for i := 0; i < len(rucksacks); i += 3 {
		triple := rucksacks[i : i+3]
		fmt.Println("\nChecking triple:", triple)

		for _, rune := range triple[0] {
			letter := string(rune)

			if strings.ContainsAny(triple[1], letter) && strings.ContainsAny(triple[2], letter) {
				fmt.Println("Found duplicate item in all rucksacks:", letter)
				prioSum += getPrio(letter)

				break
			}
		}
	}

	fmt.Printf("\n\nPriority sum of the badges is %d\n\n", prioSum)
	return prioSum
}

func main() {
	partOneResult := prioritizeDuplicates(input)
	partTwoResult := prioritizeBadges(input)

	fmt.Println("Result of part one is", partOneResult)
	fmt.Println("Result of part two is", partTwoResult)
}
