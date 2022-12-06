package main

import (
	"fmt"
	"strings"
)

func solve(data string, wantedCharAmount int) int {
	for i := wantedCharAmount; i < len(data); i++ {
		chunk := data[i-wantedCharAmount : i]
		consecutiveUniqueChars := 0
		fmt.Println("processing chunk", i, chunk)

		for _, rune := range chunk {
			occurancesInChunk := strings.Count(chunk, string(rune))
			if occurancesInChunk > 1 {
				break
			}
			consecutiveUniqueChars++
		}

		if consecutiveUniqueChars == wantedCharAmount {
			return i
		}
	}
	return -1
}

func main() {
	partOneResult := solve(input, 4)
	partTwoResult := solve(input, 14)

	fmt.Println("Result of part one is", partOneResult)
	fmt.Println("Result of part two is", partTwoResult)
}
