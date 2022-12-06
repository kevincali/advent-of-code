package main

import (
	"testing"
)

func TestPartOneTableDriven(t *testing.T) {
	var tests = []struct {
		data     string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {
			actual := solve(test.data, 4)

			if actual != test.expected {
				t.Errorf("got %d, expected %d", actual, test.expected)
			}
		})
	}
}

func TestPartTwoTableDriven(t *testing.T) {
	var tests = []struct {
		data     string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {
			actual := solve(test.data, 14)

			if actual != test.expected {
				t.Errorf("got %d, expected %d", actual, test.expected)
			}
		})
	}
}
