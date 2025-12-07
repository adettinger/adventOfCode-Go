package day7

import (
	"fmt"
	"strings"

	"github.com/adettinger/adventOfCode-Go/utils"
)

const fileName = "2025/day7/input.txt"
const sampleFileName = "2025/day7/sampleInput.txt"

func ProcessFile() {
	fileToProcess := fileName
	fmt.Printf("Processing file %q...", fileToProcess)
	lines := utils.ReadFileToSlice(fileToProcess)

	// replace S with | on first line
	lines[0] = strings.Replace(lines[0], "S", "|", -1)
	fmt.Println("Replaced S")

	totalSplits := 0
	for i := 0; i < len(lines)-1; i++ {
		// fmt.Printf("Processing line %d...\n", i)
		currSplits := 0
		lines[i+1], currSplits = processLine(lines[i], lines[i+1])
		fmt.Printf("%q Found %d splicts\n", lines[i], currSplits)
		totalSplits += currSplits
	}

	fmt.Printf("Found %d splits\n", totalSplits)
}

// Assuming there are not 2 ^ in a row
func processLine(currentLine, nextLine string) (string, int) {
	// Get position of all beams on current line
	beams := findBeams(currentLine)
	countSplits := 0
	for _, i := range beams {
		if nextLine[i:i+1] == "^" {
			countSplits++
			if i != 0 {
				nextLine = replaceCharInString(nextLine, "|", i-1)
			}
			if i != len(nextLine)-1 {
				nextLine = replaceCharInString(nextLine, "|", i+1)
			}
		} else {
			nextLine = replaceCharInString(nextLine, "|", i)
		}
	}
	return nextLine, countSplits
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
