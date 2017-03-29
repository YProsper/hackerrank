package exercices

// MergeSort sort an int slice/array with the merge algorithm.
func MergeSort(a []int, low, hi int) {
	if low < hi {
		mid := (low + hi) / 2
		MergeSort(a, low, mid)
		MergeSort(a, mid+1, hi)
		merge(a, low, mid, hi)
	}
}

func merge(a []int, low, mid, hi int) {
	ll := mid - low + 1
	rl := hi - mid

	left := make([]int, ll)
	right := make([]int, rl)

	for i := 0; i < ll; i++ {
		left[i] = a[low+i]
	}

	for i := 0; i < rl; i++ {
		right[i] = a[mid+1+i]
	}

	i, j, k := 0, 0, low
	for i < ll && j < rl {
		if left[i] <= right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
		}
		k++
	}

	if i >= ll {
		for j < rl {
			a[k] = right[j]
			j++
			k++
		}
	} else {
		for i < ll {
			a[k] = left[i]
			i++
			k++
		}
	}
}
