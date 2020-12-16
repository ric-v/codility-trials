package permcheck

import "testing"

func TestPermChecker(t *testing.T) {

	entries := []struct {
		input  []int
		result int
	}{
		{[]int{4, 1, 3, 2}, 1},
		{[]int{4, 1, 3}, 0},
		{[]int{1, 1}, 0},
	}

	for _, entry := range entries {

		result := PermChecker(entry.input)

		// check if result and expected result are same
		if entry.result != result {
			t.Errorf("PermChecker for %d failed, expected result %d, got, %d", entry.input, entry.result, result)
		}
	}
}
