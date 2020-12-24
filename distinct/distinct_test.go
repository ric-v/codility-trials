package distinct

import "testing"

func TestDistinct(t *testing.T) {

	entries := []struct {
		input  []int
		result int
	}{
		{[]int{1, 3, 5, 6, 1, 2, 3, 2, 4, 5, 5, 1}, 6},
	}

	for _, entry := range entries {

		result := Distinct(entry.input)

		// check if result and expected result are same
		if entry.result != result {
			t.Errorf("Distinct for %d failed, expected result %d, got, %d", entry.input, entry.result, result)
		}
	}
}
