package day9

import (
	"testing"

	"github.com/adettinger/adventOfCode-Go/testutils"
)

func TestDoesVerticalLineIntersecRect(t *testing.T) {
	cases := []struct {
		name   string
		l      line
		r      rectangle
		result bool
	}{
		// Defined areas: above, top, middle, bottom, below
		// For rect (0,0), (3,3), above: {}
		// Line is assumed to be sorted (lower point, higher point)
		// Out of bounds right or left
		{
			"line to left",
			line{point{-1, 0}, point{-1, 2}},
			rectangle{point{0, 0}, point{2, 2}, 4},
			false,
		},
		{
			"line to right",
			line{point{5, 0}, point{5, 2}},
			rectangle{point{0, 0}, point{2, 2}, 4},
			false,
		},
		{
			"left edge",
			line{point{0, 0}, point{0, 2}},
			rectangle{point{0, 0}, point{2, 2}, 4},
			false,
		},
		{
			"right edge",
			line{point{2, 0}, point{2, 2}},
			rectangle{point{0, 0}, point{2, 2}, 4},
			false,
		},
		// on top edge
		{
			"top  to above",
			line{point{1, 3}, point{1, 4}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			false,
		},
		{
			"middle to top",
			line{point{1, 2}, point{1, 3}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"bottom to top",
			line{point{1, 0}, point{1, 3}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"below to top",
			line{point{1, -1}, point{1, 3}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		// On bottom
		{
			"bottom to above",
			line{point{1, 0}, point{1, 4}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"bottom to middle",
			line{point{1, 0}, point{1, 2}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"below to bottom",
			line{point{1, -1}, point{1, 0}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			false,
		},
		// Other cases (only above, middle and below)
		{
			"above to above",
			line{point{1, 4}, point{1, 5}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			false,
		},
		{
			"middle to above",
			line{point{1, 2}, point{1, 4}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"below to above",
			line{point{1, -1}, point{1, 4}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"middle to middle",
			line{point{1, 1}, point{1, 2}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			false, // In the middle must have edge before and after than intersect
		},
		{
			"below to middle",
			line{point{1, -1}, point{1, 2}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"below to below",
			line{point{1, -2}, point{1, -1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			false,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if !isLineVertical(tt.l) {
				t.Error("Invalid test case: Assertion that line is vertical is false")
			}
			if !areLinesEqual(tt.l, sortOrthagonalLineAsc(tt.l)) {
				t.Error("Invalid test case: Assertion that line is sorted was false")
			}

			output := doesVerticalLineIntersectRect(tt.l, tt.r)
			testutils.AssertBool(t, output, tt.result)
		})
	}
}

func TestDoesHorizontalLineIntersecRect(t *testing.T) {
	cases := []struct {
		name   string
		l      line
		r      rectangle
		result bool
	}{
		{
			"line above",
			line{point{0, 5}, point{2, 5}},
			rectangle{point{0, 0}, point{2, 2}, 4},
			false,
		},
		{
			"line below",
			line{point{0, -1}, point{2, -1}},
			rectangle{point{0, 0}, point{2, 2}, 4},
			false,
		},
		{
			"top edge",
			line{point{0, 2}, point{2, 2}},
			rectangle{point{0, 0}, point{2, 2}, 4},
			false,
		},
		{
			"bottom edge",
			line{point{0, -0}, point{2, 0}},
			rectangle{point{0, 0}, point{2, 2}, 4},
			false,
		},
		// On left edge
		{
			"left to left edge",
			line{point{-1, 1}, point{0, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			false,
		},
		{
			"left edge to middle",
			line{point{0, 1}, point{1, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"left edge to right edge",
			line{point{0, 1}, point{3, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"left edge to right",
			line{point{0, 1}, point{4, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		// right edge
		{
			"left to right edge",
			line{point{-1, 1}, point{3, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"middle to right edge",
			line{point{1, 1}, point{3, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"right edge to right",
			line{point{3, 1}, point{4, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			false,
		},
		// Other (left, middle, right)
		{
			"left to left",
			line{point{-2, 1}, point{-1, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			false,
		},
		{
			"left to middle",
			line{point{-2, 1}, point{1, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"left to right",
			line{point{-2, 1}, point{4, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"middle to middle",
			line{point{1, 1}, point{2, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			false,
		},
		{
			"middle to right",
			line{point{1, 1}, point{4, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			true,
		},
		{
			"right to right",
			line{point{4, 1}, point{5, 1}},
			rectangle{point{0, 0}, point{3, 3}, 9},
			false,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if !isLineHorizontal(tt.l) {
				t.Error("Invalid test case: Assertion that line is horizontal is false")
			}
			if !areLinesEqual(tt.l, sortOrthagonalLineAsc(tt.l)) {
				t.Error("Invalid test case: Assertion that line is sorted was false")
			}
			output := doesHorizontalLineIntersecRect(tt.l, tt.r)
			testutils.AssertBool(t, output, tt.result)
		})
	}
}

func TestCreatePolygonFromPoints(t *testing.T) {
	cases := []struct {
		name   string
		input  []point
		result polygon
	}{
		{
			"test1",
			[]point{{1, 1}, {2, 2}, {3, 3}},
			polygon{
				[]point{{1, 1}, {2, 2}, {3, 3}},
				[]line{
					{point{1, 1}, point{3, 3}},
					{point{2, 2}, point{1, 1}},
					{point{3, 3}, point{2, 2}},
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			output, _ := createPolygonFromPoints(tt.input)
			if !arePolygonsEqual(output, tt.result) {
				t.Errorf("Got %v\nWant %v", output.String(), tt.result.String())
			}
		})
	}
}

func areSlicesOfPointsEqual(a, b []point) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if !arePointsEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func areLinesEqual(a, b line) bool {
	if !arePointsEqual(a.a, b.a) || !arePointsEqual(a.b, b.b) {
		return false
	}
	return true
}

func areSlicesOfLinesEqual(a, b []line) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if !areLinesEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func arePolygonsEqual(a, b polygon) bool {
	if !areSlicesOfPointsEqual(a.vertexes, b.vertexes) || !areSlicesOfLinesEqual(a.edges, b.edges) {
		return false
	}
	return true
}

// func TestCreatePolygonFromPoints(t *testing.T) {
// 	cases := []struct{} {}
// 	for _, tt := range cases {
// 		t.Run("", func(t *testing.T) {})
// 	}
// }
