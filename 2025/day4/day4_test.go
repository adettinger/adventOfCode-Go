package day4

import "testing"

func TestFindAccessibleRolls(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		result int
	}{
		{"basic", []string{"...", ".@.", "..."}, 1},
		{
			"sampleData",
			[]string{
				"..@@.@@@@.",
				"@@@.@.@.@@",
				"@@@@@.@.@@",
				"@.@@@@..@.",
				"@@.@@@@.@@",
				".@@@@@@@.@",
				".@.@.@.@@@",
				"@.@@@.@@@@",
				".@@@@@@@@.",
				"@.@.@@@.@.",
			},
			13,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := findAccessibleRolls(tt.input)
			if got != tt.result {
				t.Errorf("got %d expected %d", got, tt.result)
			}
		})
	}
}
