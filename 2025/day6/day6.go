package day6

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/adettinger/adventOfCode-Go/utils"
)

const fileName = "2025/day6/input.txt"
const sampleFileName = "2025/day6/sampleInput.txt"

func ProcessFilePart2() {
	fileToProcess := fileName
	fmt.Printf("Processing file %q...", fileToProcess)
	lines := utils.ReadFileToSlice(fileToProcess)
	numberLines := lines[:len(lines)-1]
	operationLine := processLine(lines[len(lines)-1])

	inProblem := false
	currentProblemNumers := make([]int, 0)
	problemSets := make([]problem, 0)
	// i is the position to process
	problemCount := 0
	for i := 0; i < maxLengthOfStrings(numberLines); i++ {
		isEmpty := true
		var sb strings.Builder
		for j, _ := range numberLines {
			if numberLines[j][i:i+1] != " " {
				isEmpty = false
				sb.WriteString(numberLines[j][i : i+1])
			}
		}
		if isEmpty {
			if inProblem {
				// Flush current numbers into a problem set
				problemSets = append(problemSets, problem{operationLine[problemCount], slices.Clone(currentProblemNumers)})
				problemCount++
				currentProblemNumers = nil
			}
			inProblem = false
		} else {
			inProblem = true
			temp, _ := strconv.Atoi(sb.String())
			currentProblemNumers = append(currentProblemNumers, temp)
		}
	}
	if inProblem {
		// Flush current numbers into a problem set
		problemSets = append(problemSets, problem{operationLine[problemCount], slices.Clone(currentProblemNumers)})
		problemCount++
	}
	for _, p := range problemSets {
		fmt.Println(p)
	}

	// Solve and total problems
	fmt.Println("Total of Solved problems: ", totalOfSolvedProblems(problemSets))
}

func maxLengthOfStrings(input []string) int {
	max := 0
	for _, i := range input {
		if max < len(i) {
			max = len(i)
		}
	}
	return max
}

func totalOfSolvedProblems(input []problem) int {
	total := 0
	for _, i := range input {
		temp, _ := i.performOperation()
		total += temp
	}
	return total
}
