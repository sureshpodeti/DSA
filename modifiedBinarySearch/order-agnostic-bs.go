package modifiedBinarySearch

/*
ar, key := []int{10, 6, 4}, 4

idx := modifiedBinarySearch.Search(ar, key)

fmt.Println(idx)
*/
func Search(ar []int, key int) int {
	left, right := 0, len(ar)-1

	isAscending := ar[left] < ar[right]

	for left <= right {
		mid := left + (right-left)/2

		if ar[mid] == key {
			return mid
		}
		if isAscending {
			if key < ar[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if key < ar[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1

}
