package day8

import (
	"fmt"
	"math"
	"strings"
)

type point struct {
	id          int
	connectedTo []int
	x           float64
	y           float64
	z           float64
}

func (p point) String() string {
	return fmt.Sprintf("{X: %d, Y: %d, Z: %d}", int(p.x), int(p.y), int(p.z))
}

func distanceBetweenPoints(a, b point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2) + math.Pow(a.z-b.z, 2))
}

type circuit struct {
	points []point
}

func (c circuit) String() string {
	var sb strings.Builder
	for _, i := range c.points {
		sb.WriteString(fmt.Sprintf("%v, ", i.String()))
	}
	output := sb.String()
	return output[:len(output)-2]
}

type pointDifference struct {
	aId      int
	bId      int
	distance float64
}

func (pd pointDifference) String() string {
	return fmt.Sprintf("a: [%d], b: [%d], distance: %f", pd.aId, pd.bId, pd.distance)
}

func printCircuits(circuits []circuit) {
	fmt.Println("Circuits:")
	for _, c := range circuits {
		fmt.Printf("%v\n", c.String())
	}
}

// func calculateCircuitsFromConnectedPoints(points []point) []circuit {
// 	// create a map to track
// 	mapPointToCircuit := make(map[int]bool, 0)
// 	circuits := make([]circuit, 0)
// 	for i, _ := range points {

// 	}

// }
