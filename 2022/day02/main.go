package main

import (
	"fmt"
	"strings"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
	win      = 6
	draw     = 3
	loss     = 0
)

func evalChoicePartOne(letter string) string {
	choice := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}
	return choice[letter]
}

func calcRoundResultPartOne(pOneChoice string, pTwoChoice string) (int, int) {
	switch pOneChoice {

	case "rock":
		switch pTwoChoice {
		case "rock":
			return (rock + draw), (rock + draw)
		case "paper":
			return (rock + loss), (paper + win)
		case "scissors":
			return (rock + win), (scissors + loss)
		}

	case "paper":
		switch pTwoChoice {
		case "rock":
			return (paper + win), (rock + loss)
		case "paper":
			return (paper + draw), (paper + draw)
		case "scissors":
			return (paper + loss), (scissors + win)
		}

	case "scissors":
		switch pTwoChoice {
		case "rock":
			return (scissors + loss), (rock + win)
		case "paper":
			return (scissors + win), (paper + loss)
		case "scissors":
			return (scissors + draw), (scissors + draw)
		}
	}

	// shouldn't happen
	return -1, -1
}

func evalChoicePartTwo(letter string) string {
	choice := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "loss",
		"Y": "draw",
		"Z": "win",
	}
	return choice[letter]
}

func calcRoundResultPartTwo(pOneChoice string, pTwoDecision string) (int, int) {
	switch pOneChoice {

	case "rock":
		switch pTwoDecision {
		case "loss":
			return (rock + win), (scissors + loss)
		case "draw":
			return (rock + draw), (rock + draw)
		case "win":
			return (rock + loss), (paper + win)
		}

	case "paper":
		switch pTwoDecision {
		case "loss":
			return (paper + win), (rock + loss)
		case "draw":
			return (paper + draw), (paper + draw)
		case "win":
			return (paper + loss), (scissors + win)
		}

	case "scissors":
		switch pTwoDecision {
		case "loss":
			return (scissors + win), (paper + loss)
		case "draw":
			return (scissors + draw), (scissors + draw)
		case "win":
			return (scissors + loss), (rock + win)
		}
	}

	// shouldn't happen
	return -1, -1
}

type evalChoice func(string) string
type calcResult func(string, string) (int, int)

func evalWinner(data string, evalChoice evalChoice, calcResult calcResult) (int, int) {
	rounds := strings.Split(data, "\n")

	var pOneSum int
	var pTwoSum int

	for i, round := range rounds {
		fmt.Printf("\n- Round %d -\n", i)
		choices := strings.Split(round, " ")

		pOneChoice := evalChoice(choices[0])
		pTwoChoice := evalChoice(choices[1])
		pOneResult, pTwoResult := calcResult(pOneChoice, pTwoChoice)
		pOneSum += pOneResult
		pTwoSum += pTwoResult

		fmt.Printf("P1: %s, P2:  %s\n", pOneChoice, pTwoChoice)
		fmt.Printf("P1 points: %d, P2 points: %d\n", pOneResult, pTwoResult)
	}

	fmt.Println("\n\n-- The game is finished! --")
	fmt.Printf("Player 1 has %d points\n", pOneSum)
	fmt.Printf("Player 2 has %d points\n", pTwoSum)

	return pOneSum, pTwoSum
}

func main() {
	evalWinner(input, evalChoicePartOne, calcRoundResultPartOne)
	evalWinner(input, evalChoicePartTwo, calcRoundResultPartTwo)
}
