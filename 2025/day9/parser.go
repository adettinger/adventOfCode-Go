package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func processFileToPoints(filename string) []point {
	fmt.Println("Processing File ", filename, "...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := make([]point, 0)
	for scanner.Scan() {
		line := scanner.Text()
		points = append(points, parseStringToPoint(line))
	}

	return points
}

// Expected format: x,y
func parseStringToPoint(input string) point {
	splitStrings := strings.Split(input, ",")
	x, _ := strconv.Atoi(splitStrings[0])
	y, _ := strconv.Atoi(splitStrings[1])
	return point{x, y}
}
