package sliding_window

/*

	ar := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 3
	sums := findSums(ar, k)
	fmt.Println(sums)
*/

func FindSums(ar []int, k int) []int {

	var result []int

	sum := 0
	start := 0

	for end := 0; end < len(ar); end++ {
		sum += ar[end]

		if end >= k-1 {
			result = append(result, sum)
			sum -= ar[start]
			start++
		}
	}

	return result
}
