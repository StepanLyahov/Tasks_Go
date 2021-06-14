package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSortConcurrencyFacade(t *testing.T) {
	arr := []int{1, 5, 4, 3, 2, 1}
	expected := []int{1, 1, 2, 3, 4, 5}

	res := MergeSortConcurrencyFacade(arr)

	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("Both arrays must be equels, but now Expected: %v, Res %v",
			expected, res)
	}
}
