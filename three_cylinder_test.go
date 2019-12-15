package cylinders

import (
	"testing"
)

func TestCounter(t *testing.T) {
	CounterBasicTesting(t, []int{2, 1}, 0)
	CounterBasicTesting(t, []int{9, 1}, 0)
	CounterBasicTesting(t, []int{5, 5, 1}, 1)
	CounterBasicTesting(t, []int{6, 10, 3, 11, 5, 2, 1}, 8)
}
func CounterBasicTesting(t *testing.T, nums []int, expected int) {
	if count := LeftHandCount(nums); count != expected {
		t.Errorf("Test Fail with input %v, expecting %v but got %v", nums, expected, count)
	}
}
func TestStringCounter(t *testing.T) {
	StringCounterBasicTesting(t, "2 1", 0)
	StringCounterBasicTesting(t, "9 1", 0)
	StringCounterBasicTesting(t, "5 5 1", 1)
	StringCounterBasicTesting(t, "6 10 3 11 5 2 1", 8)
}
func StringCounterBasicTesting(t *testing.T, nums string, expected int) {
	if count, _ := LeftHandCountString(nums); count != expected {
		t.Errorf("Test Fail with input %v, expecting %v but got %v", nums, expected, count)
	}
}
