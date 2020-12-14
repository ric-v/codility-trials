package TapeEquilibrium

import "testing"

func TestTapeEquilibrium(t *testing.T) {

	entries := []struct {
		input  []int
		result int
	}{
		{[]int{3, 1, 2, 4, 3}, 1},
		{[]int{1, 1}, 0},
		{[]int{-3, 5, -2, 1, 0, -10}, 3},
	}

	for _, entry := range entries {

		result := TapeEquilibrium(entry.input)

		// check if result and expected result are same
		if entry.result != result {
			t.Errorf("TapeEquilibrium for  %d failed, expected result %d, got, %d", entry.input, entry.result, result)
		}
	}
}
