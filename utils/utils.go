package utils

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func ReadFileToSlice(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		// fmt.Println(line)
		lines = append(lines, line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %s", err)
	}

	// Now 'lines' contains all lines from the file
	fmt.Printf("Read %d lines from the file\n", len(lines))
	return lines
}

func ProductOfInts(input []int) int {
	result := 1
	for _, i := range input {
		result = result * i
	}
	return result
}

func GetSmallestInt(input []int) (minIndex int, min int) {
	min = math.MaxInt
	for index, i := range input {
		if i < min {
			min = i
			minIndex = index
		}
	}
	return
}

func ReplaceSmallestInt(input []int, new int) []int {
	indexToReplace, _ := GetSmallestInt(input)
	input[indexToReplace] = new
	return input
}

func AbsInt(input int) int {
	if input < 0 {
		return input * -1
	}
	return input
}
