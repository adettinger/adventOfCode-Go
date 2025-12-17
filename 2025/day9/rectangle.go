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

func (r rectangle) String() string {
	return fmt.Sprintf("a: %v, b: %v, size: %d", r.a.String(), r.b.String(), r.size)
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
// Assume line is sorted
func doesEdgeIntersectsRect(l line, r rectangle) bool {
	r = sortRect(r)
	l = sortOrthagonalLineAsc(l)
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
	// top of line is top of rect
	if r.b.y == l.b.y {
		return true
	}
	// top of line is bottom of rect
	if r.a.y == l.b.y {
		return false
	}

	// bottom of line is bottom of rect
	if r.a.y == l.a.y {
		return true
	}
	// top of rect is bottom of line
	if r.b.y == l.a.y {
		return false
	}

	// crosses top line
	if l.b.y > r.b.y && l.a.y < r.b.y {
		return true
	}
	// Crosses bottom line
	if l.a.y < r.a.y && l.b.y > r.a.y {
		return true
	}
	return false
}

// Assume rect is sorted
func doesHorizontalLineIntersecRect(l line, r rectangle) bool {
	vertLine := line{point{l.a.y, l.a.x}, point{l.b.y, l.b.x}}
	flippedRect := rectangle{point{r.a.y, r.a.x}, point{r.b.y, r.b.x}, r.size}
	return doesVerticalLineIntersectRect(vertLine, flippedRect)
}
