package main

import (
	"reflect"
	"testing"
)

func TestCyclicRotation(t *testing.T) {

	entries := []struct {
		inputArr []int
		rotation int
		result   []int
	}{
		{[]int{1, 2, 3}, 1, []int{3, 1, 2}},
		{[]int{1, 2, 3}, 5, []int{2, 3, 1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
	}

	for _, entry := range entries {

		result := cyclicRotation(entry.inputArr, entry.rotation)

		// check if result and expected result are same
		if !reflect.DeepEqual(entry.result, result) {
			t.Errorf("Cyclic rotation for input array %d failed with rotation %d, expected result %d, got, %d", entry.inputArr, entry.rotation, entry.result, result)
		}
	}
}
