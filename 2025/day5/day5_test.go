package day5

import (
	"slices"
	"testing"

	testutils "github.com/adettinger/adventOfCode-Go/testutils"
)

func TestCreateRangeFromString(t *testing.T) {
	cases := []struct {
		input         string
		result        Range
		expectedError error
	}{
		{"1-2", Range{1, 2}, nil},
	}
	for _, tt := range cases {
		t.Run(tt.result.String(), func(t *testing.T) {
			output, err := createRangeFromString(tt.input)
			if tt.expectedError != nil {
				testutils.AssertError(t, err, tt.expectedError)
			} else {
				testutils.AssertNoError(t, err)
			}
			assertEqualRanges(t, output, tt.result)
		})
	}
}

func TestMergeOverlappingRanges(t *testing.T) {
	cases := []struct {
		name   string
		input  []Range
		output []Range
	}{
		{"ignores non overlap", []Range{{1, 2}, {3, 4}}, []Range{{1, 2}, {3, 4}}},
		{"merges overlap", []Range{{1, 3}, {3, 4}}, []Range{{1, 4}}},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			output := mergeOverlappedRanges(tt.input)
			if !slices.Equal(tt.output, output) {
				t.Errorf("Got %q, expected %q", output, tt.output)
			}
		})
	}
}

func TestGetValidValuesInRange(t *testing.T) {
	cases := []struct {
		name string
	}{}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func TestIsValueInRange(t *testing.T) {
	cases := []struct {
		name      string
		ranges    []Range
		value     int
		isInRange bool
	}{
		{"in range", []Range{{1, 3}}, 2, true},
		{"in range group", []Range{{1, 3}, {5, 7}}, 7, true},
		{"out of range", []Range{{1, 3}}, 7, false},
		{"out of range group", []Range{{1, 3}, {5, 7}}, 9, false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			output := isValueInRanges(tt.ranges, tt.value)
			if output != tt.isInRange {
				t.Errorf("Got %t, expected %t", output, tt.isInRange)
			}
		})
	}
}

func assertEqualRanges(t *testing.T, a, b Range) {
	t.Helper()
	if a.min != b.min || a.max != b.max {
		t.Fatalf("Ranges are not equal. %q, %q", a, b)
	}
}
