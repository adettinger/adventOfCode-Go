package day3

import "testing"

func TestFindLargestJoltage(t *testing.T) {
	cases := []struct {
		input        string
		output       string
		numBatteries int
	}{
		{"987654321111111", "987654321111", 12},
		{"811111111111119", "811111111119", 12},
		{"234234234234278", "434234234278", 12},
		{"818181911112111", "888911112111", 12},
		{"81119", "8119", 4},
	}
	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			output, err := findLargestJoltage(tt.input, tt.numBatteries)
			if err != nil {
				t.Fatal(err)
			}
			if output != tt.output {
				t.Errorf("Got %q, expected %q", output, tt.output)
			}
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
