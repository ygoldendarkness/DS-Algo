package merge_sort

func MergeSort(data []int) []int {
	if len(data) == 1 {
		return []int{data[0]}
	}
	if len(data) == 0 {
		return []int{}
	}
	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	return merge(left, right)
}

func merge(first, second []int) []int {
	result := make([]int, len(first)+len(second))
	index := 0
	firstIndex := 0
	secondIndex := 0
	for index < len(result) {
		if firstIndex == len(first) {
			result[index] = second[secondIndex]
			secondIndex++
		} else if secondIndex == len(second) {
			result[index] = first[firstIndex]
			firstIndex++
		} else if first[firstIndex] > second[secondIndex] {
			result[index] = second[secondIndex]
			secondIndex++
		} else {
			result[index] = first[firstIndex]
			firstIndex++
		}
		index++
	}
	return result
}
