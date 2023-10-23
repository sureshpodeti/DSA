package subsets

/*
ar := []int{1, 3}
ar := []int{1, 5, 3}
result := subsets.FindSubsets(ar)
fmt.Println(result)
*/
func FindSubsets(nums []int) [][]int {
	var subsets [][]int

	subsets = append(subsets, []int{})

	for _, n := range nums {
		nSubsets := len(subsets)
		for i := 0; i < nSubsets; i++ {
			var set []int
			set = append(set, subsets[i]...)
			set = append(set, n)
			subsets = append(subsets, set)
		}
	}
	return subsets
}
