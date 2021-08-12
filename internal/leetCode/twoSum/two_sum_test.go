package twoSum

import "testing"

func TestTwoSumTest(t *testing.T) {
	arr := []int{3, 3}
	target := 6

	res := twoSum(arr, target)
	expected := []int{0, 1}

	if !equal(expected, res) {
		t.Fatalf("Expected %v, but %v", expected, res)
	}
}

func equal(exp, res []int) bool {
	if len(exp) != len(res) {
		return false
	}

	for i, val := range exp {
		if val != res[i] {
			return false
		}
	}
	return true
}
