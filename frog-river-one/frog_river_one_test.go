package frogriverone

import "testing"

func TestFrogRiverOne(t *testing.T) {

	entries := []struct {
		A      []int
		x      int
		result int
	}{
		// {[]int{1, 3, 1, 4, 2, 3, 5, 4}, 5, 6},
		// {[]int{1, 2, 2, 1, 1, 1, 1}, 3, -1},
		// {[]int{1}, 2, -1},
		{[]int{1, 3, 5, 2, 4}, 0, -1},
	}

	for _, entry := range entries {

		result := FrogRiverOne(entry.A, entry.x)

		// check if result and expected result are same
		if entry.result != result {
			t.Errorf("FrogRiverOne for  %d failed, expected result %d, got, %d", entry.A, entry.result, result)
		}
	}
}
