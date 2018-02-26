package quick_sort

func partition(a []int, low, high int) int {
	pivot := a[high]

	for low < high {
		for a[low] < pivot {
			low++
		}

		for a[high] > pivot {
			high--
		}

		// Swap
		temp := a[low]
		a[low] = a[high]
		a[high] = temp
	}

	return low
}

func qSort(a []int, low int, high int) {
	if low < high {
		pivot := partition(a, low, high)
		qSort(a, low, pivot-1)
		qSort(a, pivot+1, high)
	}
}

func Sort(a []int) {
	l := 0
	h := len(a) - 1

	qSort(a, l, h)
}
