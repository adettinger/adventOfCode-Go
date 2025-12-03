package day3

import "testing"

func TestFindLargestJoltage(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}
	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			output, err := findLargestJoltage(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			assertInt(t, output, tt.output, "joltage")
		})
	}
}

func TestFindLargestDigit(t *testing.T) {
	cases := []struct {
		input           string
		largestPosition int
		largestDigit    int
	}{
		{"1234", 3, 4},
		{"98765432111111", 0, 9},
		{"81111111111111", 0, 8},
		{"23423423423427", 13, 7},
		{"81818191111211", 6, 9},
		{"12121212", 1, 2},
	}
	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			position, digit, err := findLargestDigit(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			assertInt(t, position, tt.largestPosition, "position")
			assertInt(t, digit, tt.largestDigit, "digit")
		})
	}
}

func assertInt(t *testing.T, got int, want int, name string) {
	t.Helper()
	if got != want {
		t.Fatalf("Got %d expected %d for %q", got, want, name)
	}
}
