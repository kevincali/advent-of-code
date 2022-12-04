package main

import (
	"fmt"
	"strconv"
	"strings"
)

type section struct {
	first int
	last  int
}

func newSection(sectionsHyphen string) section {
	sections := strings.Split(sectionsHyphen, "-")
	first, _ := strconv.Atoi(sections[0])
	last, _ := strconv.Atoi(sections[1])
	return section{first, last}
}

type assignment struct {
	one section
	two section
}

func newAssignment(assignmentStr string) assignment {
	sectionsHyphen := strings.Split(assignmentStr, ",")
	one := newSection(sectionsHyphen[0])
	two := newSection(sectionsHyphen[1])
	return assignment{one, two}
}

func parse(data string) []assignment {
	assignmentsStr := strings.Fields(data)

	var assignments []assignment

	for _, assignmentStr := range assignmentsStr {
		assignments = append(assignments, newAssignment(assignmentStr))
	}

	return assignments
}

func (s section) contains(other section) bool {
	return s.first >= other.first && s.last <= other.last
}

func partOne(data string) int {
	assignments := parse(data)

	var sum int
	for _, assignment := range assignments {
		sectionOne, sectionTwo := assignment.one, assignment.two
		if sectionOne.contains(sectionTwo) || sectionTwo.contains(sectionOne) {
			sum++
		}

	}
	return sum
}

func (s section) overlaps(other section) bool {
	return s.last >= other.first && s.last <= other.last || s.first >= other.first && s.first <= other.last
}

func partTwo(data string) int {
	assignments := parse(data)

	var sum int
	for _, assignment := range assignments {

		sectionOne, sectionTwo := assignment.one, assignment.two
		if sectionOne.overlaps(sectionTwo) || sectionTwo.overlaps(sectionOne) {
			sum++
		}
	}

	return sum
}

func main() {
	partOneResult := partOne(input)
	partTwoResult := partTwo(input)

	fmt.Println("Result of part one is", partOneResult)
	fmt.Println("Result of part two is", partTwoResult)
}
