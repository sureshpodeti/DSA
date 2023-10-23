package two_pointers

/*
//ar := []int{-2, -1, 0, 2, 3}
	ar := []int{-3, -1, 0, 1, 2}
	result := two_pointers.SquareSortedArray(ar)
	fmt.Println(result)
*/

func SquareSortedArray(ar []int) []int {
	start, end := 0, len(ar)-1
	updateIndex := len(ar) - 1
	newArr := make([]int, len(ar))

	for start <= end {
		left, right := ar[start]*ar[start], ar[end]*ar[end]
		if left > right {
			newArr[updateIndex] = left
			start++
		} else {
			newArr[updateIndex] = right
			end--
		}
		updateIndex--
	}
	return newArr
}
