package main

import (
	"testing"
)

func Test(t *testing.T) {
	data = ``
	actual := partOne(data)
	expected := 2

	if actual != expected {
		t.Errorf("actual != expected")
	}
}

func TestTable(t *testing.T) {
	var tests = []struct {
		data     string
		expected int
	}{
		{"", 1},
		{"", 1},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {

			actual := partOne(test.data)

			if actual != test.expected {
				t.Errorf("got %d, expected %d", actual, test.expected)
			}
		})
	}
}
