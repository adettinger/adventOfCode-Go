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
		count, err := findLargestJoltage(line)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Input: %q\n", line)
		fmt.Printf("Joltage: %d\n", count)
		total += count
	}
	fmt.Printf("Total is: %d\n", total)
}

func findLargestJoltage(line string) (int, error) {
	//Find largest digit before the last digit
	largestPosition, largestDigit, err := findLargestDigit(line[:len(line)-1])
	if err != nil {
		return -1, err
	}
	//Find the next largest digit after the previous found
	_, nextLargestDigit, err := findLargestDigit(line[largestPosition+1:])

	return (10 * largestDigit) + nextLargestDigit, nil
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
