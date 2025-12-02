package day2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"
)

var providedSample = []string{
	"11-22",
	"95-115",
	"998-1012",
	"1188511880-1188511890",
	"222220-222224",
	"1698522-1698528",
	"446443-446449",
	"38593856-38593862",
	"565653-565659",
	"824824821-824824827",
	"2121212118-2121212124",
}

func TestTotalSillyIds(t *testing.T) {
	cases := []struct {
		input  []string
		result int
	}{
		{[]string{}, 0},
		{[]string{"11-22"}, 33},
		{[]string{"998-1012"}, 1010},
		{[]string{"11-22", "998-1012"}, 1043},
	}
	for _, tt := range cases {
		t.Run(strings.Join(tt.input, ","), func(t *testing.T) {
			result, err := totalSillyIds(tt.input)
			if err != nil {
				t.Fatalf("failed with err: %q", err)
			}
			if result != tt.result {
				t.Errorf("Got %d, expected %d", result, tt.result)
			}

		})
	}
}

func TestFindSilyIds(t *testing.T) {
	cases := []struct {
		input  Range
		output []int
	}{
		{Range{11, 22}, []int{11, 22}},
		{Range{95, 115}, []int{99}},
		{Range{998, 1012}, []int{1010}},
		{Range{1188511880, 1188511890}, []int{1188511885}},
		{Range{222220, 222224}, []int{222222}},
		{Range{1698522, 1698528}, []int{}},
		{Range{446443, 446449}, []int{446446}},
		{Range{38593856, 38593862}, []int{38593859}},
		{Range{565653, 565659}, []int{}},
		{Range{824824821, 824824827}, []int{}},
		{Range{2121212118, 2121212124}, []int{}},
	}
	for _, tt := range cases {
		t.Run(getTestNameFromRanges(tt.input), func(t *testing.T) {
			result, err := findSillyIds(tt.input)
			if err != nil {
				t.Fatalf("error found: %q", err)
			}
			if !slices.Equal(result, tt.output) {
				t.Errorf("got %q expected %q from input %q", result, tt.output, tt.input)
			}
		})
	}
}

func TestIsSillyId(t *testing.T) {
	cases := []struct {
		input   int
		isSilly bool
	}{
		{11, true},
		{1010, true},
		{1188511885, true},
		{222222, true},
		{446446, true},
		{38593859, true},
		{565653, false},
		{824824821, false},
		{2121212118, false},
	}
	for _, tt := range cases {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			result := isSillyId(tt.input)
			if result != tt.isSilly {
				t.Errorf("Expected %t, got %t for input %d", tt.isSilly, result, tt.input)
			}
		})
	}
}

func getTestNameFromRanges(r Range) string {
	return fmt.Sprintf("Range from %d to %d", r.min, r.max)
}
