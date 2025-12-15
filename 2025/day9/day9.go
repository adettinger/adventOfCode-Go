package day9

import (
	"errors"
	"fmt"
	"os"
)

// Assume rects are sorted
func getMaxRectSizeInPoly(poly polygon, sortedRects []rectangle) (rectangle, error) {
	for _, i := range sortedRects {
		fmt.Printf("Testing (%v, %v, %d)\n", i.a.String(), i.b.String(), i.size)
		if isRectInPoly(i, poly) {
			return i, nil
		}
	}
	return rectangle{}, errors.New("No rectagles contained in polygon")
}

const fileName = "2025/day9/input.txt"
const sampleFileName = "2025/day9/sampleInput.txt"

func Day9() {
	fileToProcess := fileName
	points := processFileToPoints(fileToProcess)
	poly, err := createPolygonFromPoints(points)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sortedRects := generateSortedListOfRectsFromPoints(points)

	maxRect, err := getMaxRectSizeInPoly(poly, sortedRects)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Found max size of: %d\nFrom points %v and %v", maxRect.size, maxRect.a.String(), maxRect.b.String())
}
