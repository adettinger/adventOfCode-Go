package day6

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/adettinger/adventOfCode-Go/testutils"
)

// func TestTotalSolvedProblems(t *testing.T) {
// 	cases := []struct{}{}
// 	for _, tt := range cases {
// 		t.Run()
// 	}
// }

func TestParseStringsToProblems(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]string
		output []problem
	}{
		{"1", [][]string{{"1", "2"}, {"3", "4"}, {"*", "+"}}, []problem{{"*", []int{1, 3}}, {"+", []int{2, 4}}}},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result := parseStringsToProblems(tt.input)
			if !areSliceOfProblemsEqual(result, tt.output) {
				t.Errorf("slices are not equal. Got %q, want %q", arrayOfProblemsToString(result), arrayOfProblemsToString(tt.output))
			}
		})
	}
}

func TestPerformOperation(t *testing.T) {
	cases := []struct {
		problem problem
		result  int
	}{
		{problem{"*", []int{2, 3, 4}}, 24},
		{problem{"+", []int{2, 3, 4}}, 9},
	}
	for _, tt := range cases {
		t.Run(tt.problem.operation, func(t *testing.T) {
			output, err := tt.problem.performOperation()
			if err != nil {
				t.Errorf("Error found performing operation")
			}
			if output != tt.result {
				testutils.AssertInts(t, output, tt.result)
			}
		})
	}
}

func TestProcessLine(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{"1 2 3", []string{"1", "2", "3"}},
		{"  1   2    3   ", []string{"1", "2", "3"}},
		{"  +  *  +", []string{"+", "*", "+"}},
	}
	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			result := processLine(tt.input)

			if !slices.Equal(result, tt.output) {
				t.Errorf("got %q, expected %q", strings.Join(result, ","), strings.Join(result, ","))
			}
		})
	}
}

func areSliceOfProblemsEqual(a, b []problem) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if !a[i].Equal(b[i]) {
			return false
		}
	}
	return true
}

func arrayOfProblemsToString(problems []problem) string {
	var sb strings.Builder
	for _, i := range problems {
		sb.WriteString(fmt.Sprintf("[%q], ", i.String()))
	}
	return sb.String()
}
