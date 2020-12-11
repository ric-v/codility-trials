package missingelement

import "testing"

func TestMissingElem(t *testing.T) {

	entries := []struct {
		input  []int
		result int
	}{
		{[]int{1, 2, 4}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16, 17}, 11},
	}

	for _, entry := range entries {

		result := MissingElement(entry.input)

		// check if result and expected result are same
		if entry.result != result {
			t.Errorf("MissingElem for  %d failed, expected result %d, got, %d", entry.input, entry.result, result)
		}
	}
}
