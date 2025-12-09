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

type rect struct {
	min point
	max point
}

type line struct {
	a point
	b point
}

type polygon struct {
	edges []line
}

func createPolygonFromPoints(points []point) polygon {
	edges := make([]line, len(points))
	edges[0] = line{points[0], points[len(points)-1]}
	for i := 1; i < len(points); i++ {
		edges[i] = line{points[i], points[i-1]}
	}
	return polygon{edges}
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
func getMaxRectSizeInPoly(poly polygon, points []point) (int, point, point) {
	maxSize := 0
	pointA := point{}
	pointB := point{}
	for i, _ := range points {
		for j := i + 1; j < len(points); j++ {
			fmt.Printf("Testing (%d, %d)", i, j)
			if isRectInPoly(rect{points[i], points[j]}, poly) {
				rectSize := getSizeOfRectangeFromPoints(points[i], points[j])
				if rectSize > maxSize {
					maxSize = rectSize
					pointA = points[i]
					pointB = points[j]
				}
			}
		}
	}
	return maxSize, pointA, pointB
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
	fileToProcess := sampleFileName
	points := processFileToPoints(fileToProcess)
	poly := createPolygonFromPoints(points)

	size, a, b := getMaxRectSizeInPoly(poly, points)
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
