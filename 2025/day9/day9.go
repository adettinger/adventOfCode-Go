package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/adettinger/adventOfCode-Go/utils"
)

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("{x: %d, y: %d}", p.x, p.y)
}

func sliceOfPointsToString(points []point) string {
	var sb strings.Builder
	for _, i := range points {
		sb.WriteString(i.String() + ", ")
	}
	output := sb.String()
	return output[:len(output)-2]
}

func getSizeOfRectangeFromPoints(a, b point) int {
	return utils.AbsInt(a.x-b.x+1) * utils.AbsInt(a.y-b.y+1)
}

func getMaxRectSize(points []point) (int, point, point) {
	maxSize := 0
	pointA := point{}
	pointB := point{}
	for i, _ := range points {
		for j := i + 1; j < len(points); j++ {
			rectSize := getSizeOfRectangeFromPoints(points[i], points[j])
			if rectSize > maxSize {
				maxSize = rectSize
				pointA = points[i]
				pointB = points[j]
			}
		}
	}
	return maxSize, pointA, pointB
}

const fileName = "2025/day9/input.txt"
const sampleFileName = "2025/day9/sampleInput.txt"

func Day9() {
	fileToProcess := fileName
	points := processFileToPoints(fileToProcess)
	fmt.Println("Found points")
	for _, i := range points {
		println(i.String())
	}

	size, a, b := getMaxRectSize(points)
	fmt.Printf("Found max size of: %d\nFrom points %v and %v", size, a.String(), b.String())
}

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
