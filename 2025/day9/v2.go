package day9

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	stack "github.com/golang-collections/collections/stack"
)

func V2() {
	fileToProcess := fileName
	points := processFileToPoints(fileToProcess)

	// Get sorted unique x and y cords
	xValues, yValues := getSortedUniqueDimensions(points)

	// map sorted lists to indecies
	xMap := map[int]int{}
	for i, x := range xValues {
		xMap[x] = i
	}

	yMap := map[int]int{}
	for i, y := range yValues {
		yMap[y] = i
	}

	// fill points with "."
	grid := make([][]string, len(yValues))
	for i, _ := range grid {
		grid[i] = make([]string, len(xValues))
		for j, _ := range grid[i] {
			grid[i][j] = "."
		}
	}
	fmt.Println("Empty")
	printGrid(grid)

	// Set corner points
	mappedPoints := make([]point, len(points))
	for i, p := range points {
		grid[yMap[p.y]][xMap[p.x]] = "#"
		mappedPoints[i] = point{x: xMap[p.x], y: yMap[p.y]}
	}
	fmt.Println("Corners")
	printGrid(grid)

	// rasterize
	for i, _ := range mappedPoints {
		l := line{a: mappedPoints[i], b: mappedPoints[(i+1)%len(mappedPoints)]}
		l = sortOrthagonalLineAsc(l)

		if isLineVertical(l) {
			for y := l.a.y; y < l.b.y; y++ {
				grid[y][l.a.x] = "#"
			}
		} else {
			for x := l.a.x; x < l.b.x; x++ {
				grid[l.a.y][x] = "#"
			}
		}
	}
	fmt.Println("Edges")
	printGrid(grid)

	// fill polygon
	startingPoint, _ := getInsidePoint(grid)
	grid = floodFill(grid, startingPoint)
	fmt.Println("Filled")
	printGrid(grid)

	// test rects
	sortedRects := generateSortedListOfRectsFromPoints(points)
	for _, r := range sortedRects {
		mappedRect := rectangle{
			a:    point{xMap[r.a.x], yMap[r.a.y]},
			b:    point{xMap[r.b.x], yMap[r.b.y]},
			size: r.size,
		}
		if isEnclosed(mappedRect, grid) {
			fmt.Printf("Found max rect: %v\n", r.String())
			os.Exit(0)
		}
	}
}

func printGrid(grid [][]string) {
	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Printf("%v\n", strings.Join(grid[i], ""))
	}
}

func getSortedUniqueDimensions(points []point) ([]int, []int) {
	encounteredX := map[int]bool{}
	encounteredY := map[int]bool{}
	xValues := make([]int, 0)
	yValues := make([]int, 0)

	for _, point := range points {
		if !encounteredX[point.x] {
			encounteredX[point.x] = true
			xValues = append(xValues, point.x)
		}

		if !encounteredY[point.y] {
			encounteredY[point.y] = true
			yValues = append(yValues, point.y)
		}
	}
	sort.Ints(xValues)
	sort.Ints(yValues)
	return xValues, yValues
}

func getInsidePoint(grid [][]string) (point, error) {
	for y, _ := range grid {
		for x, _ := range grid[y] {
			if grid[y][x] == "." {
				continue
			}

			hitsLeft := 0
			prev := "."

			for i := x; i >= 0; i-- {
				cur := grid[y][i]
				if cur != prev {
					hitsLeft++
				}
				prev = cur
			}
			if hitsLeft%2 == 1 {
				return point{x, y}, nil
			}
		}
	}
	return point{}, errors.New("No inside point found")
}

// Doesnt work in cases where the first point is also filled (small grids)
func floodFill(grid [][]string, start point) [][]string {
	s := stack.New()
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	s.Push(start)

	for s.Len() > 0 {
		p, _ := s.Pop().(point)
		if grid[p.y][p.x] != "." {
			continue
		}
		grid[p.y][p.x] = "X"
		for _, dir := range directions {
			nx := p.x + dir[0]
			ny := p.y + dir[1]
			if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[0]) {
				if grid[ny][nx] == "." {
					s.Push(point{nx, ny})
				}
			}
		}
	}
	return grid
}

func isEnclosed(r rectangle, grid [][]string) bool {
	r = sortRect(r)
	for x := r.a.x; x <= r.b.x; x++ {
		if grid[r.a.y][x] == "." || grid[r.b.y][x] == "." {
			return false
		}
	}

	for y := r.a.y; y <= r.b.y; y++ {
		if grid[y][r.a.x] == "." || grid[y][r.b.x] == "." {
			return false
		}
	}
	return true
}
