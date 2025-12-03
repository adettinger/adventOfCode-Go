package day3

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const fileName = "2025/day3/input.txt"
const sampleFileName = "2025/day3/sampleInput.txt"
const numberOfBatteries = 12

func ProcessFile() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		countString, err := findLargestJoltage(line, numberOfBatteries)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Input: %q\n", line)
		fmt.Printf("Joltage: %q\n", countString)
		count, err := strconv.Atoi(countString)
		if err != nil {
			log.Fatal(err)
		}
		total += count
	}
	fmt.Printf("Total is: %d\n", total)
}

func findLargestJoltage(line string, batteriesRemaining int) (string, error) {
	if len(line) <= batteriesRemaining {
		return line, nil
	}

	//Find largest digit before the last digit
	largestPosition, _, err := findLargestDigit(line[:len(line)-batteriesRemaining+1])
	if err != nil {
		return "", err
	}
	if batteriesRemaining >= 2 {
		subPart, err := findLargestJoltage(line[largestPosition+1:], batteriesRemaining-1)
		if err != nil {
			log.Fatal(err)
		}
		return line[largestPosition:largestPosition+1] + subPart, nil
	} else {
		return line[largestPosition : largestPosition+1], nil
	}
}

// @return index, value
func findLargestDigit(input string) (int, int, error) {
	largestPosition := 0
	largestDigit := 0

	for i := 0; i < len(input); i++ {
		charInt, err := strconv.Atoi(input[i : i+1])
		if err != nil {
			return -1, -1, errors.New("Error converting string to int")
		}
		if charInt > largestDigit {
			largestDigit = charInt
			largestPosition = i
		}
	}

	return largestPosition, largestDigit, nil
}
