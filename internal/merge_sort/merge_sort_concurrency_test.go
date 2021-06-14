package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeFuncSizesEqual(t *testing.T) {
	left := []int{1, 3, 5}
	right := []int{2, 4, 6}

	expected := []int{1, 2, 3, 4, 5, 6}

	res := merge(left, right)

	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("Both arrays must be equels, but now Expected: %v, Res %v",
			expected, res)
	}
}

func TestMergeFuncSizesDifferent(t *testing.T) {
	left := []int{5}
	right := []int{2, 4, 6, 7}

	expected := []int{2, 4, 5, 6, 7}

	res := merge(left, right)

	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("Both arrays must be equels, but now Expected: %v, Res %v",
			expected, res)
	}
}

func TestMergeFuncSizesOne(t *testing.T) {
	left := []int{5}
	right := []int{4}

	expected := []int{4, 5}

	res := merge(left, right)

	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("Both arrays must be equels, but now Expected: %v, Res %v",
			expected, res)
	}
}

func TestMergeSortConcurrency(t *testing.T) {
	arr := []int{1, 5, 4, 3, 2, 1}
	expected := []int{1, 1, 2, 3, 4, 5}

	res := MergeSortConcurrency(arr)

	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("Both arrays must be equels, but now Expected: %v, Res %v",
			expected, res)
	}

}

func TestMergeSortConcurrencyZeroArr(t *testing.T) {
	arr := make([]int, 0)
	expected := make([]int, 0)

	res := MergeSortConcurrency(arr)

	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("Both arrays must be equels, but now Expected: %v, Res %v",
			expected, res)
	}

}
