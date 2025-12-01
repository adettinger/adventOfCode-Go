package day1

import (
	"testing"
)

func TestProcessMoves(t *testing.T) {
	processMovesTests := []struct {
		name            string
		moves           []string
		expectedPasses  int
		expectedLands   int
		isErrorExpected bool
	}{
		{"Basic", []string{"L50"}, 1, 1, false},
		{"Basic2", []string{"L1"}, 0, 0, false},
		{
			"Provided sample",
			[]string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"},
			6,
			3,
			false,
		},
	}

	for _, tt := range processMovesTests {
		t.Run(tt.name, func(t *testing.T) {
			passes, lands, err := ProcessMoves(tt.moves)

			if tt.isErrorExpected {
				if err == nil {
					t.Errorf("Expected to get an error but did not")
				}
			} else {
				if err != nil {
					t.Errorf("Expected not to get an error but did. Error %q", err.Error())
				}
				if passes != tt.expectedPasses {
					t.Errorf("Passes: Expected to get %d but got %d", tt.expectedPasses, passes)
				}
				if lands != tt.expectedLands {
					t.Errorf("Lands: Expected to get %d but got %d", tt.expectedLands, lands)
				}
			}

		})
	}
}

func TestNextPosition(t *testing.T) {
	nextPositionTests := []struct {
		name             string
		currentPosition  int
		move             string
		expectedPosition int
		expectedPasses   int
		isErrorExpected  bool
	}{
		{"right from 1", 1, "R2", 3, 0, false},
		{"left 2 from 1", 1, "L2", 99, 1, false},
		{"left 102 from 1", 1, "L102", 99, 2, false},
		{"right 2 from 99", 99, "R2", 1, 1, false},
		{"right 102 from 99", 99, "R102", 1, 2, false},
		{"right 50 from 50", 50, "L50", 0, 1, false},
		{"left 50 from 50", 50, "R50", 0, 1, false},
		{"left 2 from 0", 0, "L2", 98, 0, false},
		{"left 102 from 0", 0, "L102", 98, 1, false},
		{"right 2 from 0", 0, "R2", 2, 0, false},
		{"right 102 from 0", 0, "R102", 2, 1, false},
		{"error on invalid letter", 1, "Q2", 1, -1, true},
		{"error on invalid number", 1, "AA", 1, -1, true},
		{"error on missing letter", 1, "1", 1, -1, true},
		{"error on missing number", 1, "A", 1, -1, true},
		{"error on empty input", 1, "", 1, -1, true},
	}

	for _, tt := range nextPositionTests {
		t.Run(tt.name, func(t *testing.T) {
			nextPosition, passes, err := NextPosition(tt.currentPosition, tt.move)

			//Check error status
			if tt.isErrorExpected {
				if err == nil {
					t.Errorf("Expected to get an error but did not for input\ncurrent position: %d\nmove: %q. Got result %d", tt.currentPosition, tt.move, nextPosition)
				}
			} else {
				if err != nil {
					t.Errorf("Expected not to get an error but did for input\ncurrent position: %d\nmove: %q. Got error %q", tt.currentPosition, tt.move, err.Error())
				}
				if nextPosition != tt.expectedPosition {
					t.Errorf("Expected to get %d but got %d for input\ncurrent position: %d\nmove: %q", tt.expectedPosition, nextPosition, tt.currentPosition, tt.move)
				}
				if passes != tt.expectedPasses {
					t.Errorf("Expected %d passes but got %d passes", tt.expectedPasses, passes)
				}
			}

		})
	}
}
