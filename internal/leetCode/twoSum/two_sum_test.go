package twoSum

import "testing"

func TestTwoSumTest(t *testing.T) {
	//TODO прикрутить табличные тесты
	arr := []int{3, 3}
	target := 6

	res := twoSum(arr, target)
	expected := []int{0, 1}

	if !equal(expected, res) {
		t.Fatalf("Expected %v, but %v", expected, res)
	}

	//arr := []int{3, 2, 4}
	//target := 6
	//
	//res := twoSum(arr, target)
	//expected := []int{2, 1}
	//
	//if !equal(expected, res) {
	//	t.Fatalf("Expected %v, but %v", expected, res)
	//}

	//arr := []int{0, 4, 3, 0}
	//target := 0
	//
	//res := twoSum(arr, target)
	//expected := []int{0, 3}
	//
	//if !equal(expected, res) {
	//	t.Fatalf("Expected %v, but %v", expected, res)
	//}
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
