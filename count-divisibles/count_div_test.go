package countdiv

import "testing"

func TestCountDiv(t *testing.T) {

	entries := []struct {
		A          int
		B          int
		K          int
		divisibles int
	}{
		{6, 11, 2, 3},
		{100, 1230000, 2, 614951},
		{10, 10, 5, 1},
	}

	for _, entry := range entries {

		result := CountDiv(entry.A, entry.B, entry.K)

		// check if result and expected result are same
		if entry.divisibles != result {
			t.Errorf("CountDiv for numbers within A %d and B %d with K %d failed, expected result %d, got, %d", entry.A, entry.B, entry.K, entry.divisibles, result)
		}
	}
}
