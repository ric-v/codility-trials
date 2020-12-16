package missingint

import (
	"testing"
)

func TestMissingInteger(t *testing.T) {

	entries := []struct {
		input  []int
		result int
	}{
		{[]int{1, 3, 6, 4, 1, 2}, 5},
		{[]int{1, 2, 3}, 4},
		{[]int{-1, -3}, 1},
	}

	for _, entry := range entries {

		result := MissingInteger(entry.input)

		// check if result and expected result are same
		if entry.result != result {
			t.Errorf("MissingInteger for %d failed, expected result %d, got, %d", entry.input, entry.result, result)
		}
	}
}
