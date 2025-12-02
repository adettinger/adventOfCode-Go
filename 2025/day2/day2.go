package day2

import (
	"errors"
	"fmt"
	"strconv"
)

type Range struct {
	min int
	max int
}

func (r Range) String() string {
	return fmt.Sprintf("Ranges{min: %d, max: %d}", r.min, r.max)
}

func ProcessFile() {
	rangeStrings, _ := parseFileToStrings()
	total, _ := totalSillyIds(rangeStrings)
	fmt.Printf("Total: %d", total)
}

func totalSillyIds(input []string) (int, error) {
	total := 0
	for _, i := range input {
		rangeTotal := 0
		r, err := parseRanges(i)
		if err != nil {
			return 0, err
		}
		fmt.Printf("Processing Range: %q\n", r)

		sillyIds, err := findSillyIds(r)
		if err != nil {
			return 0, err
		}

		for _, j := range sillyIds {
			fmt.Printf("Found sillyId: %d\n", j)
			rangeTotal += j
		}
		fmt.Printf("Range total: %d\n", rangeTotal)
		total += rangeTotal
	}
	return total, nil
}

func findSillyIds(r Range) ([]int, error) {
	sillyIds := []int{}

	if r.min > r.max {
		return []int{}, errors.New("max is smaller than min")
	}

	for i := r.min; i <= r.max; i++ {
		if isSillyId(i) {
			sillyIds = append(sillyIds, i)
		}
	}

	return sillyIds, nil
}

func isSillyId(input int) bool {
	str := strconv.Itoa(input)
	for i := 1; i <= (len(str) / 2); i++ {
		if areAllSubStringsEqual(str, i) {
			return true
		}
	}
	return false
}

func areAllSubStringsEqual(input string, size int) bool {
	if len(input)%size != 0 || len(input) <= size {
		return false
	}

	for i := 0; i < len(input); i += size {
		end := i + size
		// low := input[:size]
		// high := input[i:end]
		// fmt.Printf("Comparing %q and %q\n", low, high)
		if input[:size] != input[i:end] {
			return false
		}
	}
	return true
}
