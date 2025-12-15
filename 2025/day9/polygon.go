package day9

import (
	"fmt"
	"strings"
)

/*
// Structs
*/
type polygon struct {
	vertexes []point
	edges    []line
}

func (poly polygon) String() string {
	var vertexesSB strings.Builder
	for _, i := range poly.vertexes {
		vertexesSB.WriteString(i.String() + ", ")
	}
	vertexString := vertexesSB.String()
	vertexString = vertexString[:len(vertexString)-2]

	var edgesSB strings.Builder
	for _, i := range poly.edges {
		edgesSB.WriteString(i.String() + ", ")
	}
	edgesString := edgesSB.String()
	edgesString = edgesString[:len(edgesString)-2]

	return fmt.Sprintf("{vertexes: %v, edges: %v}", vertexString, edgesString)
}

func createPolygonFromPoints(points []point) (polygon, error) {
	prevPoint := point{0, 0}
	prev2Point := point{0, 0}

	edges := make([]line, len(points))
	edges[0] = line{points[0], points[len(points)-1]}
	for i := 1; i < len(points); i++ {
		edge := line{points[i], points[i-1]}
		if !isLineOrthagonal(edge) {
			return polygon{}, fmt.Errorf("Line is not orthagonal: %v", edge.String())
		}

		if arePointsEqual(prevPoint, points[i]) {
			return polygon{}, fmt.Errorf("prev point is equal at index: %d", i)
		}
		if arePointsEqual(prev2Point, points[i]) {
			return polygon{}, fmt.Errorf("prev2 point is equal at index: %d", i)
		}
		prev2Point = prevPoint
		prevPoint = points[i]

		edges[i] = line{points[i], points[i-1]}
	}
	return polygon{vertexes: points, edges: edges}, nil
}

/*
// Functions
*/

func isRectInPoly(r rectangle, poly polygon) bool {
	// TODO: Check a corner just inside is in the rectangle

	for _, edge := range poly.edges {
		if doesEdgeIntersectsRect(edge, r) {
			return false
		}
	}
	return true
}

// Implement ray-crossing algo
// Count how many times poly intersects ray horizontally left from point
/*
func isPointInPolygon(poly polygon, p point) bool {
	countIntersections := 0
	for _, edge := range poly.edges {
		// TODO: Check if point is on the edge


		intersects, _ := doesLineIntersectLeftRayV2(edge, p)
		if intersects {
			countIntersections++
		}
	}
	if countIntersections%2 == 0 {
		return false
	}
	return true
}

// Ray through point with -1 slope
// Since lines are edges of polygon, they are not parallel to ray
// Inconsistant if point is on the line (does not count as intersecting)
func doesLineIntersectLeftRayV2(edge line, p point) (bool, error) {
	ray := createSlopeInterceptFromPointSlope(-1.0, p)
	lineToTest := createSlopeInterceptFromLine(edge)
	Xval, err := calculateXvalFromIntersectionOfLines(ray, lineToTest)
	if err != nil {
		return false, fmt.Errorf("doesLineIntersectLeftRayV2: Found parallel lines. \nEdge: %v\nPoint: %v", edge, p)
	}
	if Xval < float64(p.x) {
		return true, nil
	}
	return false, nil
}


// TODO: Change to ray at an angle
func doesLineIntersectLeftRayV1(edge line, p point) bool {
	if edge.a.y != edge.b.y { //line is not horizontal
		if edge.a.x < p.x { //line must be left of point
			minY, maxY := utils.SortInts(edge.a.y, edge.b.y)
			if minY < p.y && maxY > p.y {
				return true
			}
		}
	}
	return false
}
*/
