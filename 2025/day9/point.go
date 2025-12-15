package day9

import (
	"fmt"
	"strings"
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

func arePointsEqual(a, b point) bool {
	return a.x == b.x && a.y == b.y
}
