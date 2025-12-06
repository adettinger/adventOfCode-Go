package day6

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type problem struct {
	operation string
	values    []int
}

func (p problem) performOperation() (int, error) {
	switch {
	case p.operation == "+":
		return sumInts(p.values), nil
	case p.operation == "*":
		return productInts(p.values), nil
	default:
		return -1, fmt.Errorf("Invalid operation: %q", p.operation)
	}
}

func (a problem) Equal(b problem) bool {
	if a.operation != b.operation {
		return false
	}
	if !slices.Equal(a.values, b.values) {
		return false
	}
	return true
}

func (p problem) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Operation: %q, Values: ", p.operation))
	for _, val := range p.values {
		sb.WriteString(fmt.Sprint(strconv.Itoa(val), ", "))
	}
	return sb.String()
}

func sumInts(input []int) int {
	total := 0
	for _, i := range input {
		total += i
	}
	return total
}

func productInts(input []int) int {
	total := 1
	for _, i := range input {
		total = total * i
	}
	return total
}