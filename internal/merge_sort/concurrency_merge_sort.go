package merge_sort

func MergeSortConcurrencyFacade(arr []int) []int {
	resChan := make(chan []int)
	defer close(resChan)
	go mergeSortConcurrency(arr, resChan)

	return <-resChan
}

func mergeSortConcurrency(arr []int, res chan []int) {
	size := len(arr)
	if size == 0 {
		res <- arr
	}

	if size == 1 {
		res <- arr
	}

	mid := size / 2

	leftChan := make(chan []int)
	rightChan := make(chan []int)

	go mergeSortConcurrency(arr[:mid], leftChan)
	go mergeSortConcurrency(arr[mid:], rightChan)

	resLeft := <-leftChan
	resRight := <-rightChan

	resArr := merge(resLeft, resRight)

	res <- resArr
}
