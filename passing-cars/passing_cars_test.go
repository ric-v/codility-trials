package passingcars

import "testing"

func TestPassingCars(t *testing.T) {

	entries := []struct {
		input  []int
		result int
	}{
		{[]int{0, 1, 0, 1, 1}, 5},
		{[]int{1, 0}, 0},
	}

	for _, entry := range entries {

		result := PassingCars(entry.input)

		// check if result and expected result are same
		if entry.result != result {
			t.Errorf("PassingCars for  %d failed, expected result %d, got, %d", entry.input, entry.result, result)
		}
	}
}
