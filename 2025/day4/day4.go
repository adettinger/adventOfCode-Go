package day4

import (
	"fmt"

	utils "github.com/adettinger/adventOfCode-Go/utils"
)

type position struct {
	row int
	col int
}

const fileName = "2025/day4/input.txt"
const sampleFileName = "2025/day4/sampleInput.txt"
const maxSurrounded = 3

func ProcessFile() {
	fileToProcess := fileName
	fmt.Printf("Processing file %q...", fileToProcess)
	input := utils.ReadFileToSlice(fileToProcess)

	total := 0
	latestFound := 1
	loopCounter := 1
	for latestFound > 0 {
		fmt.Println("Loop: ", loopCounter)
		latestFound, input = findAndRemoveAccessableRolls(input)
		total += latestFound
		loopCounter++
	}
	fmt.Printf("Total Found: %d\n", total)
}

func printFoundIndexes(found []position) {
	fmt.Print("Found indexes: ")
	for _, i := range found {
		fmt.Printf("[%d, %d],  ", i.row, i.col)
	}
	fmt.Println()
}

func findAndRemoveAccessableRolls(input []string) (int, []string) {
	total, found := findAccessibleRolls(input)
	fmt.Println("Found to be removed: ", total)
	printFoundIndexes(found)
	input = removeFoundFromState(input, found)
	return total, input
}

func printState(input []string) {
	for _, i := range input {
		fmt.Println(i)
	}
}

func removeFoundFromState(state []string, found []position) []string {
	for _, pos := range found {
		state[pos.row] = state[pos.row][:pos.col] + "." + state[pos.row][pos.col+1:]
	}
	return state
}

func findAccessibleRolls(input []string) (int, []position) {
	total := 0
	indexes := []position{}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if isIndexPaper(input, i, j) == 1 {
				surroundedCount := 0
				hasPreviousRow := i != 0
				hasNextrow := i+1 < len(input)
				hasPreviousCol := j != 0
				hasNextCol := j+1 < len(input[i])

				if hasPreviousRow {
					if hasPreviousCol {
						surroundedCount += isIndexPaper(input, i-1, j-1)
					}
					surroundedCount += isIndexPaper(input, i-1, j)
					if hasNextCol {
						surroundedCount += isIndexPaper(input, i-1, j+1)
					}
				}
				if hasPreviousCol {
					surroundedCount += isIndexPaper(input, i, j-1)
				}
				if hasNextCol {
					surroundedCount += isIndexPaper(input, i, j+1)
				}
				if hasNextrow {
					if hasPreviousCol {
						surroundedCount += isIndexPaper(input, i+1, j-1)
					}
					surroundedCount += isIndexPaper(input, i+1, j)
					if hasNextCol {
						surroundedCount += isIndexPaper(input, i+1, j+1)
					}
				}
				// fmt.Printf("Paper at %d, %d has %d neighbors\n", i, j, surroundedCount)

				if surroundedCount <= maxSurrounded {
					indexes = append(indexes, position{i, j})
					total++
				}
			}
		}
	}
	return total, indexes
}

func isIndexPaper(input []string, i, j int) int {
	if input[i][j:j+1] == "@" {
		return 1
	}
	return 0
}
