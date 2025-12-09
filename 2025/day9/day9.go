package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

type solvedRect struct {
	min  point
	max  point
	size int
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
func getMaxRectSizeInPoly(poly polygon, rects []solvedRect) solvedRect {
	// sort rects
	sort.Slice(rects, func(i int, j int) bool {
		return rects[i].size > rects[j].size
	})

	for _, i := range rects {
		fmt.Printf("Testing (%v, %v, %d)\n", i.min.String(), i.max.String(), i.size)
		if isRectInPoly(i, poly) {
			return i
		}
	}
	return solvedRect{}
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

	sortedRects := generateSortedListOfRects(points)

	maxRect := getMaxRectSizeInPoly(poly, sortedRects)
	fmt.Printf("Found max size of: %d\nFrom points %v and %v", maxRect.size, maxRect.min.String(), maxRect.max.String())
}

func generateSortedListOfRects(points []point) []solvedRect {
	sortedRects := make([]solvedRect, 0)
	for i, _ := range points {
		for j := i + 1; j < len(points); j++ {
			rectSize := getSizeOfRectangeFromPoints(points[i], points[j])
			sortedRects = append(sortedRects, solvedRect{points[i], points[j], rectSize})
		}
	}
	return sortedRects
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
