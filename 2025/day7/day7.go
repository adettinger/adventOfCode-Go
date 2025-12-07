package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adettinger/adventOfCode-Go/utils"
)

const fileName = "2025/day7/input.txt"
const sampleFileName = "2025/day7/sampleInput.txt"

type row struct {
	line   string
	counts []int
}

func createRow(line string) row {
	return row{line, make([]int, len(line))}
}

func (r row) String() string {
	var sb strings.Builder
	for _, i := range r.counts {
		sb.WriteString(strconv.Itoa(i) + ",")
	}
	return fmt.Sprintf("%s %s", r.line, sb.String())
}

func ProcessFile() {
	fileToProcess := fileName
	fmt.Printf("Processing file %q...", fileToProcess)
	lines := utils.ReadFileToSlice(fileToProcess)

	// replace S with | on first line
	// lines[0] = strings.Replace(lines[0], "S", "|", -1)

	rows := convertInputToRows(lines)
	rows[0] = convertFirstLine(rows[0])

	totalSplits := 0
	for i := 0; i < len(rows)-1; i++ {
		// fmt.Printf("Processing line %d...\n", i)
		currSplits := 0
		rows[i+1], currSplits = processLine(rows[i], rows[i+1])
		fmt.Printf("%s Found %d splicts\n", rows[i].String(), currSplits)
		totalSplits += currSplits
	}

	fmt.Printf("Found %d splits\n", totalSplits)

	fmt.Printf("Found %d permutations\n", sumInts(rows[len(rows)-1].counts))
}

func convertFirstLine(input row) row {
	for i, _ := range input.line {
		if input.line[i:i+1] == "S" {
			input.line = replaceCharInString(input.line, "|", i)
			input.counts[i] = 1
		}
	}
	return input
}

func convertInputToRows(input []string) []row {
	result := make([]row, len(input))
	for index, i := range input {
		result[index] = createRow(i)
	}
	return result
}

// Assuming there are not 2 ^ in a row
func processLine(currentRow, nextRow row) (row, int) {
	// Get position of all beams on current line
	beams := findBeams(currentRow.line)
	countSplits := 0
	for _, i := range beams {
		currBeamQuant := currentRow.counts[i]
		if nextRow.line[i:i+1] == "^" {
			countSplits++
			if i != 0 {
				nextRow.line = replaceCharInString(nextRow.line, "|", i-1)
				nextRow.counts[i-1] += currBeamQuant
			}
			if i != len(nextRow.line)-1 {
				nextRow.line = replaceCharInString(nextRow.line, "|", i+1)
				nextRow.counts[i+1] += currBeamQuant
			}
		} else {
			nextRow.line = replaceCharInString(nextRow.line, "|", i)
			nextRow.counts[i] += currBeamQuant
		}
	}
	return nextRow, countSplits
}

func replaceCharInString(input, char string, index int) string {
	return input[:index] + char + input[index+1:]
}

func findBeams(line string) []int {
	beams := make([]int, 0)
	i := 0
	for i < len(line) {
		foundIndex := strings.Index(line[i:], "|")
		if foundIndex < 0 {
			return beams
		}
		beams = append(beams, i+foundIndex)
		i += foundIndex + 1
	}
	return beams
}

func sumInts(input []int) int {
	total := 0
	for _, i := range input {
		total += i
	}
	return total
}
