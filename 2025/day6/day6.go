package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "2025/day6/input.txt"
const sampleFileName = "2025/day6/sampleInput.txt"

func ProcessFile() {
	fileToProcess := fileName
	fmt.Println("Processing File ", fileToProcess, "...")
	file, err := os.Open(fileToProcess)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		processedLine := processLine(line)
		lines = append(lines, processedLine)
	}

	problems := parseStringsToProblems(lines)
	fmt.Printf("Total of problems: %d", totalSolvesProblems(problems))
}

func totalSolvesProblems(input []problem) int {
	total := 0
	for _, i := range input {
		temp, _ := i.performOperation()
		total += temp
	}
	return total
}

func parseStringsToProblems(input [][]string) []problem {
	problems := make([]problem, len(input[0]))
	for i := 0; i < len(input[0]); i++ {
		values := make([]int, len(input)-1)
		for j := 0; j < len(input)-1; j++ {
			values[j], _ = strconv.Atoi(input[j][i])
		}

		problems[i] = problem{operation: input[len(input)-1][i], values: values}
	}
	return problems
}

func processLine(input string) []string {
	splitStrings := strings.Split(input, " ")

	result := make([]string, 0)
	for _, i := range splitStrings {
		if i != "" && i != " " {
			result = append(result, i)
		}
	}
	return result
}
