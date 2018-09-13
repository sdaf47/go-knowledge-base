package sort

func MergeSort(m []int) (result []int) {
	if len(m) <= 1 {
		return m
	}

	middle := len(m) / 2
	left := make([]int, middle)
	right := make([]int, len(m))

	left = m[:middle]
	right = m[middle:]

	left = MergeSort(left)
	right = MergeSort(right)

	result = merge(left, right)

	return result
}

func merge(a, b []int) (result []int) {
	la := 0
	lb := 0
	for la < len(a) && lb < len(b) {
		if a[la] < b[lb] {
			result = append(result, a[la])
			la++
		} else {
			result = append(result, b[lb])
			lb++
		}
	}
	for la < len(a) {
		result = append(result, a[la])
		la++
	}
	for lb < len(b) {
		result = append(result, b[lb])
		lb++
	}
	return
}
