package day8

import (
	"fmt"
	"math"
	"testing"
)

var samplePoints = []point{
	{0.0, 0.0, 0.0},
	{1.0, 1.0, 1.0},
	{2.0, 2.0, 2.0},
}

func TestParseStringToPoint(t *testing.T) {
	cases := []struct {
		input  string
		result point
	}{
		{"1,2,3", point{1, 2, 3}},
	}
	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			output := parseStringToPoint(tt.input)
			assertPointsAreEqual(t, output, tt.result)
		})
	}
}

func TestDistanceBetweenPoints(t *testing.T) {
	cases := []struct {
		input1 point
		input2 point
		result float64
	}{
		{point{0, 0, 0}, point{1, 1, 1}, math.Sqrt(3.0)},
	}
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v - %v", tt.input1.String(), tt.input2.String()), func(t *testing.T) {
			output := distanceBetweenPoints(tt.input1, tt.input2)
			if !areFloatsApproxEqual(output, tt.result) {
				t.Errorf("got %f want %f", output, tt.result)
			}
		})
	}
}

func TestInsertToSortedListOfDistance(t *testing.T) {
	cases := []struct {
		name     string
		base     []pointDifference
		toInsert pointDifference
		result   []pointDifference
	}{
		{
			"test 1",
			[]pointDifference{
				{a: samplePoints[0], b: samplePoints[0], distance: 0.0},
				{a: samplePoints[0], b: samplePoints[2], distance: math.Sqrt(12)},
			},
			pointDifference{a: samplePoints[0], b: samplePoints[1], distance: math.Sqrt(3)},
			[]pointDifference{
				{a: samplePoints[0], b: samplePoints[0], distance: 0.0},
				{a: samplePoints[0], b: samplePoints[1], distance: math.Sqrt(3)},
				{a: samplePoints[0], b: samplePoints[2], distance: math.Sqrt(12)},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			output := insertToSortedListOfDistances(tt.base, tt.toInsert)
			if !areSliceOfDifferencesEqual(output, tt.result) {
				t.Errorf("got %v, want %v", output, tt.result)
			}
		})
	}
}

// func TestGetSmallestInt(t *testing.T) {
// 	cases := []struct{
// 		input []int
// 		resultIndex int
// 		resultInt int
// 	} {}
// 	for _, tt := range cases {
// 		t.Run()
// 	}
// }

// func TestParseStringToPoint(t *testing.T) {
// 	cases := []struct{
// 	} {}
// 	for _, tt := range cases {
// 		t.Run()
// 	}
// }

// func TestParseStringToPoint(t *testing.T) {
// 	cases := []struct{
// 	} {}
// 	for _, tt := range cases {
// 		t.Run()
// 	}
// }

func areFloatsApproxEqual(a, b float64) bool {
	const totolerance = 1e-7
	return math.Abs(a-b) < totolerance
}

func arePointsEqual(a, b point) bool {
	if a.x != b.x || a.y != b.y || a.z != b.z {
		return false
	}
	return true
}

func assertPointsAreEqual(t *testing.T, got, want point) {
	t.Helper()
	if !arePointsEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func areCircuitsEqual(t *testing.T, got, want circuit) bool {
	t.Helper()
	if len(got.points) != len(want.points) {
		return false
	}
	for i, _ := range got.points {
		if !arePointsEqual(got.points[i], want.points[i]) {
			return false
		}
	}
	return true
}

func arePointDifferencesEqual(got, want pointDifference) bool {
	if !arePointsEqual(got.a, want.a) || !arePointsEqual(got.b, want.b) || !areFloatsApproxEqual(got.distance, want.distance) {
		return false
	}
	return true
}

func areSliceOfDifferencesEqual(got, want []pointDifference) bool {
	if len(got) != len(want) {
		return false
	}
	for i, _ := range got {
		if !arePointDifferencesEqual(got[i], want[i]) {
			return false
		}
	}
	return true
}
