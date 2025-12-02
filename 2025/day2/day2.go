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
	if len(str)%2 != 0 {
		return false
	}
	if str[:len(str)/2] == str[len(str)/2:] {
		return true
	}
	return false
}
