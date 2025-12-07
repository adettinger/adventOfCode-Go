package day7

import (
	"slices"
	"testing"

	"github.com/adettinger/adventOfCode-Go/testutils"
)

func TestReplaceCharInString(t *testing.T) {
	cases := []struct {
		input  string
		char   string
		index  int
		output string
	}{
		{"aaa", "b", 1, "aba"},
	}
	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			output := replaceCharInString(tt.input, tt.char, tt.index)
			testutils.AssertStrings(t, output, tt.output)
		})
	}
}

func TestFindBeams(t *testing.T) {
	cases := []struct {
		input  string
		output []int
	}{
		{"...............", []int{}},
		{".......|.......", []int{7}},
		{".|.....|...|...", []int{1, 7, 11}},
	}
	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			output := findBeams(tt.input)
			if !slices.Equal(output, tt.output) {
				t.Errorf("got %v, expected %v", output, tt.output)
			}
		})
	}
}
