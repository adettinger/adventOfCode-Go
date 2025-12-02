package day2

import "testing"

func TestParseRanges(t *testing.T) {
	cases := []struct {
		input  string
		result Range
	}{
		{"11-22", Range{11, 22}},
		{"95-115", Range{95, 115}},
		{"998-1012", Range{998, 1012}},
		{"1188511880-1188511890", Range{1188511880, 1188511890}},
		{"222220-222224", Range{222220, 222224}},
		{"1698522-1698528", Range{1698522, 1698528}},
		{"446443-446449", Range{446443, 446449}},
		{"38593856-38593862", Range{38593856, 38593862}},
		{"565653-565659", Range{565653, 565659}},
		{"824824821-824824827", Range{824824821, 824824827}},
		{"2121212118-2121212124", Range{2121212118, 2121212124}},
	}
	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			result, err := parseRanges(tt.input)
			if err != nil {
				t.Fatalf("error found: %q", err)
			}
			if result != tt.result {
				t.Errorf("got %q expected %q from input %q", result, tt.result, tt.input)
			}
		})
	}
}
