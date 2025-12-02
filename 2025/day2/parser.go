package day2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	utils "github.com/adettinger/adventOfCode-Go/utils"
)

const fileName = "2025/day2/input.txt"
const sampleFileName = "2025/day2/sampleInput.txt"

func parseFileToStrings() ([]string, error) {
	fileStrings := utils.ReadFileToSlice(fileName)
	if len(fileStrings) != 1 {
		return []string{}, errors.New("File has more than 1 line")
	}
	rangeStrings := strings.Split(fileStrings[0], ",")
	return rangeStrings, nil
}

// Separate string "123-123" into "123", "123"
func parseRanges(input string) (Range, error) {
	splitStrings := strings.Split(input, "-")
	if len(splitStrings) != 2 {
		return Range{}, fmt.Errorf(
			"Expected split strings to have 2 parts, found %d \ninput: %q",
			len(splitStrings),
			input,
		)
	}
	min, err := strconv.Atoi(splitStrings[0])
	max, err2 := strconv.Atoi(splitStrings[1])
	if err != nil || err2 != nil {
		return Range{}, fmt.Errorf(
			"Error parsing strings to int\nerrror1: %q\nerror2:%q",
			err,
			err2,
		)
	}

	return Range{min, max}, nil
}
