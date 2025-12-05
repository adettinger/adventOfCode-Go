package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const fileName = "2025/day5/input.txt"
const sampleFileName = "2025/day5/sampleInput.txt"

type Range struct {
	min int
	max int
}

func (r Range) String() string {
	return fmt.Sprintf("[min: %d, max: %d]", r.min, r.max)
}

func sliceOfRangesToString(ranges []Range) string {
	var sb strings.Builder

	for _, r := range ranges {
		sb.WriteString(r.String() + ",")
	}
	return sb.String()
}

// TODO: Convert to using linked list or binary tree not slices
func ProcessFile() {
	fileToProcess := fileName
	fmt.Println("Processing File ", fileToProcess, "...")
	file, err := os.Open(fileToProcess)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse Ranges
	ranges := make([]Range, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("Found break")
			break
		}

		fmt.Println("Found line: ", line)
		convRange, _ := createRangeFromString(line)
		ranges = append(ranges, convRange)
	}
	fmt.Println("Found ranges: ", ranges)

	//SortRanges
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})
	fmt.Println("Sorted ranges: ", ranges)

	ranges = mergeOverlappedRanges(ranges)
	fmt.Println("Merged ranges: ", ranges)

	// Parse inputs
	inputs := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parsedInt, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Could not parse input %q to int. Error: %q", line, err)
		}
		inputs = append(inputs, parsedInt)
	}
	fmt.Println("Found inputs: ", inputs)

	// Part 1
	validInputs := getValidValuesInRanges(ranges, inputs)
	fmt.Println("ValidInputs: ", validInputs)
	fmt.Printf("Found %d valid Inputs\n", len(validInputs))

	// Part 2
	possibleTotal := getCountOfPossibleInputs(ranges)
	fmt.Println("Count of possible valid inputs: ", possibleTotal)
}

func createRangeFromString(input string) (Range, error) {
	splitStrings := strings.Split(input, "-")
	if len(splitStrings) != 2 {
		return Range{}, fmt.Errorf("Could not split %q correctly", input)
	}
	min, err := strconv.Atoi(splitStrings[0])
	max, err2 := strconv.Atoi(splitStrings[1])
	if err != nil || err2 != nil {
		return Range{}, fmt.Errorf("Could not convert split strings (%q, %q) to ints. Found error: %q, %q", splitStrings[0], splitStrings[1], err, err2)
	}
	return Range{min: min, max: max}, nil
}

//Assume ranges are sorted
func mergeOverlappedRanges(input []Range) []Range {
	index := 0
	for index < len(input)-1 {
		if input[index].max >= input[index+1].min {
			input[index].max = maxInt(input[index].max, input[index+1].max)
			input = append(input[:index+1], input[index+2:]...)
		} else {
			index++
		}
	}
	return input
}

func getValidValuesInRanges(ranges []Range, values []int) []int {
	output := make([]int, 0)
	for _, i := range values {
		if isValueInRanges(ranges, i) {
			output = append(output, i)
		}
	}
	return output
}

// TODO: Assume sorted ranges, use binary search
func isValueInRanges(ranges []Range, value int) bool {
	l := 0
	r := len(ranges) - 1

	for l <= r {
		index := (r + l) / 2
		if ranges[index].min <= value {
			if ranges[index].max >= value {
				return true
			} else { //range is to the left
				l = index + 1
			}
		} else { //range is to the right
			r = index - 1
		}
	}
	return false
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Assume ranges are sorted and de-conflicted
func getCountOfPossibleInputs(ranges []Range) int {
	total := 0
	for _, r := range ranges {
		total += r.max - r.min + 1
	}
	return total
}
