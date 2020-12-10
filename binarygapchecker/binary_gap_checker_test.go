package binarygapchecker

import (
	"testing"
)

func TestCyclicRotation(t *testing.T) {

	entries := []struct {
		input  int
		result int
	}{
		{1, 0},
		{41, 2},
		{12032, 1},
		{300120, 3},
	}

	for _, entry := range entries {

		result := BinaryGapChecker(entry.input)

		// check if result and expected result are same
		if entry.result != result {
			t.Errorf("Binary gap for input %d returned wrong value than expected, expected result %d, got, %d", entry.input, entry.result, result)
		}
	}
}
