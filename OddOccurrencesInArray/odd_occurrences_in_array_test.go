package oddoccurrencesinarray

import (
	"testing"
)

func TestCyclicRotation(t *testing.T) {

	entries := []struct {
		inputArr []int
		result   int
	}{
		{[]int{1, 2, 1, 3, 3}, 2},
		{[]int{2, 3, 1, 2, 3}, 1},
		{[]int{1, 2, 2, 1, 3}, 3},
	}

	for _, entry := range entries {

		result := OddOccurrencesInArray(entry.inputArr)

		// check if result and expected result are same
		if result != entry.result {
			t.Errorf("Cyclic rotation for input array %d failed, expected result %d, got, %d", entry.inputArr, entry.result, result)
		}
	}
}
