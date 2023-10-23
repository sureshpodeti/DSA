package cyclic_sort

/*
	//ar := []int{3, 1, 5, 4, 2}
	//ar := []int{2, 6, 4, 3, 1, 5}
	//ar := []int{1, 5, 6, 4, 3, 2} <- NOT WORKING CHECK AGAIN
	ar := []int{1, 2, 3}
	cyclic_sort.CycleSort(ar)

	fmt.Println(ar)
*/

func swap(ar []int, i, j int) {
	t := ar[i]
	ar[i], ar[j] = ar[j], t
}
func CycleSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		j := ar[i] - 1
		if ar[i] != (i + 1) {
			//swap
			swap(ar, i, j)
		}
	}
}
