package day9

import (
	"errors"
	"fmt"
	"math"

	"github.com/adettinger/adventOfCode-Go/utils"
)

/*
// Structs
*/
type line struct {
	a point
	b point
}

func (l line) String() string {
	return fmt.Sprintf("{a: %v, b: %v}", l.a.String(), l.b.String())
}

// y = mx + b
type slopeInterceptLine struct {
	m float64
	b float64
}

func (sil slopeInterceptLine) String() string {
	return fmt.Sprintf("{m: %f, b: %f}", sil.m, sil.b)
}

func createSlopeInterceptFromPointSlope(m float64, p point) slopeInterceptLine {
	b := float64(p.y) - (m * float64(p.x))
	return slopeInterceptLine{m, b}
}

func createSlopeInterceptFromLine(input line) slopeInterceptLine {
	m := 0.0
	// if line is vertical, slope is infinite
	if input.a.x-input.b.x == 0 {
		m = math.Inf(1)
	} else {
		m = (float64(input.a.y - input.b.y)) / (float64(input.a.x - input.b.x))
	}
	return createSlopeInterceptFromPointSlope(m, input.a)
}

/*
// Functions
*/

func calculateXvalFromIntersectionOfLines(a, b slopeInterceptLine) (float64, error) {
	// lines are parallel
	if a.m == b.m {
		return 0.0, errors.New("Lines are parallel")
	}
	return (a.b - b.b) / (b.m - a.m), nil
}

func isLineOrthagonal(l line) bool {
	if isLineVertical(l) || isLineHorizontal(l) {
		return true
	}
	return false
}

func isLineVertical(l line) bool {
	if l.a.x == l.b.x {
		return true
	}
	return false
}

func isLineHorizontal(l line) bool {
	if l.a.y == l.b.y {
		return true
	}
	return false
}

func sortOrthagonalLineAsc(l line) line {
	if isLineVertical(l) {
		l.a.y, l.b.y = utils.SortIntsAsc(l.a.y, l.b.y)
		return l
	} else {
		l.a.x, l.b.x = utils.SortIntsAsc(l.a.x, l.b.x)
		return l
	}
}
