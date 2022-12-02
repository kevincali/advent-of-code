package main

import (
	"testing"
)

func TestEvalWinnerDraw(t *testing.T) {
	var data = `A Y
B X
C Z`

	pOneSum, pTwoSum := evalWinner(data, evalChoicePartOne, calcRoundResultPartOne)

	if pOneSum != pTwoSum {
		t.Errorf("the game should be a draw")
	}
}

func TestEvalWinnerPermutations(t *testing.T) {
	var data = `A X
A Y
A Z
B X
B Y
B Z
C X
C Y
C Z`

	pOneSum, pTwoSum := evalWinner(data, evalChoicePartOne, calcRoundResultPartOne)

	if pOneSum != pTwoSum {
		t.Errorf("the game should be a draw")
	}
}

func TestEvalWinnerWin(t *testing.T) {
	var data = `A Y`

	pOneSum, pTwoSum := evalWinner(data, evalChoicePartOne, calcRoundResultPartOne)

	if pOneSum > pTwoSum {
		t.Errorf("player two should win")
	}
}
