package main

import (
	"fmt"
	"strconv"
	"strings"
)

type stack []string

type operation struct {
  amount int
  src int
  dst int
}

type supply struct {
  stacks []stack
  operations []operation
}

func newOperation(operationStr string) operation {
  fields := strings.Fields(operationStr)
  amount, _:= strconv.Atoi(fields[1])
  src, _:= strconv.Atoi(fields[3])
  dst, _:= strconv.Atoi(fields[5])
  return operation{amount, src-1, dst-1}
}

func (s *stack) push(crate string) {
  *s = append(*s, crate)
}

func (s *stack) pop() string {
	if len(*s) == 0 {
		return ""
	} 
  last := len(*s)-1 
  crate := (*s)[last] 
  *s = (*s)[:last] 
  return crate
}

func newStacks(stacksStr []string) []stack {
  lastLineStr := stacksStr[len(stacksStr)-1]
  lastLine := strings.Fields(lastLineStr)
  amount, _ := strconv.Atoi(lastLine[len(lastLine)-1])
  
  stacks := make([]stack, amount)
  for i:=len(stacksStr)-2; i >= 0; i-- {
    currentStack := stacksStr[i]
  
    charIndex := 1
    for j:=0; j<amount; j++ {
      crate := string(currentStack[charIndex])
      // only grab actual crates
      if crate != " " {
        stacks[j].push(crate)
      }
      charIndex += 4 
    }
  }
  return stacks
}

func parse(data string) supply {
  parts := strings.Split(data, "\n\n")
  stacksStr := strings.Split(parts[0], "\n")
  stacks := newStacks(stacksStr)
  
  operationsStr := strings.Split(parts[1], "\n")
  var operations []operation
  for _, operationStr := range operationsStr {
    operations = append(operations, newOperation(operationStr))
  }
  
  return supply{stacks, operations}
}

func partOne(data string) string {
  supply := parse(data)
  
  for i, operation := range supply.operations {
    fmt.Printf("Executing operation %d: move %d from stack %d to stack %d\n", i+1, operation.amount, operation.src, operation.dst)

    for j:=0; j<operation.amount; j++ {
      crate := supply.stacks[operation.src].pop()
      supply.stacks[operation.dst].push(crate)
    }
  }
  
  var topCrates []string 
  for _, stack := range supply.stacks {
    topCrates = append(topCrates, stack[len(stack)-1])
  }
  
	return strings.Join(topCrates, "")
}

func partTwo(data string) string {
  supply := parse(data)

  for i, operation := range supply.operations {
    fmt.Printf("Executing operation %d: move %d from stack %d to stack %d\n", i+1, operation.amount, operation.src, operation.dst)
    
    var tmpCrates []string
    for j:=0; j<operation.amount; j++ {
      // use pop and prepend to tmp array to preserve order
      crate := supply.stacks[operation.src].pop()
      tmpCrates = append([]string{crate}, tmpCrates...)
    }

    // add crates in order
    for _, crate := range tmpCrates {
      supply.stacks[operation.dst].push(crate)
    }
  }

  var topCrates []string 
  for _, stack := range supply.stacks {
    topCrates = append(topCrates, stack[len(stack)-1])
  }

	return strings.Join(topCrates, "")
}

func main() {
	partOneResult := partOne(input)
	partTwoResult := partTwo(input)

	fmt.Println("Result of part one is", partOneResult)
	fmt.Println("Result of part two is", partTwoResult)
}
