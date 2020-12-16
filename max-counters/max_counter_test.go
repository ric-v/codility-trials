package maxcounters

import (
	"reflect"
	"testing"
)

func TestMaxCounter(t *testing.T) {

	entries := []struct {
		input  []int
		N      int
		result []int
	}{
		{[]int{3, 4, 4, 6, 1, 4, 4}, 5, []int{3, 2, 2, 4, 2}},
		{[]int{4, 4, 4, 4, 4, 4, 4, 4}, 5, []int{0, 0, 0, 8, 0}},
	}

	for _, entry := range entries {

		result := MaxCounter(entry.input, entry.N)

		// check if result and expected result are same
		if !reflect.DeepEqual(entry.result, result) {
			t.Errorf("MaxCounter for  %d failed for N %d, expected result %d, got, %d", entry.input, entry.N, entry.result, result)
		}
	}
}
