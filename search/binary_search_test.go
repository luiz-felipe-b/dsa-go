package search

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 6, 7, 8, 9}
	t.Run("should return -1 if target is not found", func(t *testing.T) {
		assertEqual(t, -1, BinarySearch(arr, 5))
	})

	t.Run("should return the target if found", func(t *testing.T) {
		assertEqual(t, 6, BinarySearch(arr, 6))
	})
}

func assertEqual(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
