package day1

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const StartingPosition = 50
const fileName = "2025/day1/input.txt"

func ReadFile() []string {
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

func ProcessMovesFromFile() {
	moves := ReadFile()
	passes, landsOnZero, err := ProcessMoves(moves)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Found Result:\nPasses: %d\nLands on Zero: %d", passes, landsOnZero)
	}
}

// @return passes, landsOnZero, error
func ProcessMoves(moves []string) (int, int, error) {
	currentPosition := StartingPosition
	passesCount := 0
	passes := 0
	landsOnZeroCount := 0
	var error error
	for index, val := range moves {
		currentPosition, passes, error = NextPosition(currentPosition, val)
		if error != nil {
			return -1, -1, fmt.Errorf("Error processing inputs at index %d, error: %q", index, error.Error())
		}
		passesCount += passes
		if currentPosition == 0 {
			landsOnZeroCount++
		}
	}
	return passesCount, landsOnZeroCount, nil
}

// Count as pass if it lands on 0
// @return newPosition, passes, error
func NextPosition(currentPosition int, move string) (int, int, error) {
	if len(move) < 2 {
		return -1, -1, fmt.Errorf("NextPosition did not receive valid move")
	}

	letter := string(move[0])
	numberStr := move[1:]
	spacesToMove, err := strconv.Atoi(numberStr)
	if err != nil {
		return -1, -1, fmt.Errorf("Error converting to number: %q", err.Error())
	}

	//Check assertions
	if spacesToMove < 0 {
		return -1, -1, errors.New("Cannot move negative positions")
	}
	if spacesToMove == 0 {
		return currentPosition, 0, nil
	}

	switch letter {
	case "L": //Subtracting position
		finalPosition := currentPosition - spacesToMove
		if finalPosition > 0 {
			return finalPosition, 0, nil
		} else {
			finalPosition = finalPosition % 100
			passes := (spacesToMove-currentPosition)/100 + 1
			if currentPosition == 0 {
				passes--
			}
			// if finalPosition == 0 {
			// 	passes += 1
			// } else
			if finalPosition < 0 {
				finalPosition += 100
			}
			return finalPosition, passes, nil
		}
	case "R": //Adding position
		finalPosition := (currentPosition + spacesToMove) % 100
		passes := (currentPosition + spacesToMove) / 100
		return finalPosition, passes, nil
	default:
		return -1, -1, fmt.Errorf("Error reading letter from move")
	}
}
