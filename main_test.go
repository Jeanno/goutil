package goutil

import (
    "testing"
)

func TestSliceEqual(t *testing.T) {

    // A list of integers
    intList := []int{1, 2, 3, 4, 5}
    // A list of integers
    intList2 := []int{1, 2, 3, 4, 5}
    intList3 := []int{1, 2, 1, 4, 5}
    // Assert that the mapped list is correct
    if !SliceEqual(intList, intList2) {
        t.Error("Expected true, got false")
    }
    if SliceEqual(intList, intList3) {
        t.Error("Expected false, got true")
    }
}

func TestSliceMap(t *testing.T) {
    // A list of integers
    intList := []int{1, 2, 3, 4, 5}
    mapped := SliceMap(intList, func(i int) int {
        return i * 2
    })
    // Assert that the mapped list is correct
    if !SliceEqual(mapped, []int{2, 4, 6, 8, 10}) {
        t.Error("Expected true, got false")
    }
}
