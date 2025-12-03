package utils

import (
	"bufio"
	"fmt"
	"log"
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
