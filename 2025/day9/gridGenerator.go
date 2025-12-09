package day9

import "github.com/adettinger/adventOfCode-Go/utils"

func isRectInPoly(r rect, poly polygon) bool {
	// check that a corner of the rectangle is in the polygon
	if !isPointInPolygon(poly, r.min) {
		return false
	}
	for _, edge := range poly.edges {
		if doesLineIntersectsRect(edge, r) {
			return false
		}
	}
	return true
}

func isPolygonRightAngledOnly(points []point) bool {
	prevPoint := point{}
	currPoint := point{}
	for i := 1; i < len(points); i++ {
		prevPoint = points[i-1]
		currPoint = points[i]
		if prevPoint.x != currPoint.x && prevPoint.y != currPoint.y {
			return false
		}
	}
	return true
}

// Implement ray-crossing algo
// Count how many times poly intersects ray horizontally left from point
func isPointInPolygon(poly polygon, p point) bool {
	countIntersections := 0
	for _, i := range poly.edges {
		if doesLineIntersectLeftRay(i, p) {
			countIntersections++
		}
	}
	if countIntersections%2 == 0 {
		return false
	}
	return true

}

func doesLineIntersectLeftRay(edge line, p point) bool {
	if edge.a.y != edge.b.y { //line is not horizontal
		if edge.a.x < p.x { //line must be left of point
			minY, maxY := sortInts(edge.a.y, edge.b.y)
			if minY < p.y && maxY > p.y {
				return true
			}
		}
	}
	return false
}

func sortInts(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// Assume line is orthagonal
func doesLineIntersectsRect(l line, r rect) bool {
	if l.a.x == l.b.x { //vertical line
		if l.a.x < r.min.x || l.a.x > r.max.x {
			return false //line is left or right of rectangle
		}
		//y positions must be both above or both below rect
		if l.a.y < r.min.y && l.b.y < r.min.y {
			return false
		} else if l.a.y > r.max.y && l.b.y > r.max.y {
			return false
		} else {
			return true
		}
	} else { // horizontal line
		// TODO!!!!!
	}
}

func isPointInOrOnRectangle(p point, r rect) bool {
	r = sortRect(r)
	if r.min.x <= p.x && r.min.y <= p.y && r.max.x >= p.x && r.max.y >= p.y {
		return true
	}
	return false
}

func sortRect(r rect) rect {
	xPositions := []int{r.min.x, r.max.x}
	yPositions := []int{r.min.y, r.max.y}

	r.min = point{utils.MinInt(xPositions), utils.MinInt(yPositions)}
	r.max = point{utils.MaxInt(xPositions), utils.MaxInt(yPositions)}

	return r
}
