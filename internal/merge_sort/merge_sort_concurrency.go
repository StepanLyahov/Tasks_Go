package merge_sort

func MergeSortConcurrency(arr []int) []int {
	size := len(arr)
	if size == 0 {
		return arr
	}

	if size == 1 {
		return arr
	}

	mid := size / 2

	leftArr := MergeSortConcurrency(arr[:mid])
	rightArr := MergeSortConcurrency(arr[mid:])

	return merge(leftArr, rightArr)
}

func merge(left, right []int) []int {
	sizeL := len(left)
	sizeR := len(right)

	resArr := make([]int, sizeL+sizeR)
	pointL, pointR, pointerRes := 0, 0, 0

	for pointerRes != len(resArr) {
		if pointL == sizeL && pointR != sizeR {
			resArr[pointerRes] = right[pointR]
			pointR++
		} else if pointR == sizeR && pointL != sizeL {
			resArr[pointerRes] = left[pointL]
			pointL++
		} else {
			if left[pointL] <= right[pointR] {
				resArr[pointerRes] = left[pointL]
				pointL++
			} else {
				resArr[pointerRes] = right[pointR]
				pointR++
			}
		}
		pointerRes++
	}

	return resArr
}
