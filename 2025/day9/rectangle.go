package day9

import (
	"fmt"
	"sort"

	"github.com/adettinger/adventOfCode-Go/utils"
)

/*
// Structs
*/

type rectangle struct {
	a    point
	b    point
	size int
}

func createRectangleFromPoints(a, b point) rectangle {
	return rectangle{a, b, getSizeOfRectangeFromPoints(a, b)}
}

func getSizeOfRectangeFromPoints(a, b point) int {
	return utils.AbsInt(a.x-b.x+1) * utils.AbsInt(a.y-b.y+1)
}

/*
// Functions
*/

// Points -> solvedRect[]
func generateSortedListOfRectsFromPoints(points []point) []rectangle {
	sortedRects := make([]rectangle, 0)

	for i, _ := range points {
		fmt.Printf("Generating rects for %v\n", points[i].String())
		for j := i + 1; j < len(points); j++ {
			sortedRects = append(sortedRects, createRectangleFromPoints(points[i], points[j]))
		}
	}

	fmt.Println("Sorting list of rects...")
	sort.Slice(sortedRects, func(i int, j int) bool {
		return sortedRects[i].size > sortedRects[j].size
	})
	fmt.Println("List sorted")
	return sortedRects
}

/*
func isPointInOrOnRectangle(p point, r rect) bool {
	r = sortRect(r)
	if r.min.x <= p.x && r.min.y <= p.y && r.max.x >= p.x && r.max.y >= p.y {
		return true
	}
	return false
}
*/

// Recalculate corners of rectangle to be (minX, minY), (maxX, maxY)
func sortRect(r rectangle) rectangle {
	xPositions := []int{r.a.x, r.b.x}
	yPositions := []int{r.a.y, r.b.y}

	r.a = point{utils.MinInt(xPositions), utils.MinInt(yPositions)}
	r.b = point{utils.MaxInt(xPositions), utils.MaxInt(yPositions)}

	return r
}

// Assume line is orthagonal
// TODO: Rethink, consider lines that overlap
// Consider if line edge is on edge of rectangle
func doesEdgeIntersectsRect(l line, r rectangle) bool {
	r = sortRect(r)
	l = sortOrthagonalLine(l)
	if isLineVertical(l) { //vertical line
		return doesVerticalLineIntersectRect(l, r)
	} else { // horizontal line
		return doesHorizontalLineIntersecRect(l, r)
	}
}

// Assume rect is sorted
func doesVerticalLineIntersectRect(l line, r rectangle) bool {
	if l.a.x <= r.a.x || l.a.x >= r.b.x {
		return false //line is left or right of rectangle
	}

	/*
		// Handle line ends on rect edge
	*/
	if r.b.y == l.b.y {
		// top of line is top of rect
		return true
	}
	if r.a.y == l.b.y {
		// top of line is bottom of rect
		return false
	}

	if r.a.y == l.a.y {
		// bottom of line is bottom of rect
		return true
	}
	if r.b.y == l.a.y {
		// top of rect is bottom of line
		return false
	}

	//TODOOOOOOOOOO: Handle line and rect dont share point but may intersect

	// TODO: Check top line, check bottom line
	//y positions must be both above or both below rect
	if l.a.y < r.a.y && l.b.y < r.a.y {
		return false
	} else if l.a.y > r.b.y && l.b.y > r.b.y {
		return false
	} else {
		return true
	}
}

// Assume rect is sorted
func doesHorizontalLineIntersecRect(l line, r rectangle) bool {
	if l.a.y <= r.a.y || l.a.y >= r.b.y {
		return false //line is above or below
	}
	// TODO: check left line, check right line
	if l.a.x < r.a.x && l.b.x < r.a.x {
		return false
	} else if l.a.x > r.b.x && l.b.x > r.b.x {
		return false
	} else {
		return true
	}
}
