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

func getMaxRectSize(points []point) int {

}

const fileName = "2025/day8/input.txt"
const sampleFileName = "2025/day8/sampleInput.txt"

func day9() {
	fileToProcess := fileName
	points := processFileToPoints(fileToProcess)
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
