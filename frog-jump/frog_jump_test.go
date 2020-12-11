package frogjmp

import "testing"

func TestCyclicRotation(t *testing.T) {

	entries := []struct {
		x      int
		y      int
		jump   int
		result int
	}{
		{10, 85, 30, 3},
		{85, 850, 10, 77},
		{19, 20, 6, 1},
		{5, 1000000000, 2, 499999998},
	}

	for _, entry := range entries {

		result := FrogJump(entry.x, entry.y, entry.jump)

		// vet the result from func
		if result != entry.result {
			t.Errorf("Frog jump for x %d, y %d with jump %d returned wrong value than expected, expected result %d, got, %d", entry.x, entry.y, entry.jump, entry.result, result)
		}
	}
}
