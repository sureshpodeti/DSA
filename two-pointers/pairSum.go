package two_pointers

type Pair struct {
	Start int
	End   int
}

func PairSum(ar []int, target int) Pair {

	start, end := 0, len(ar)-1

	for start < end {
		sum := ar[start] + ar[end]

		if sum == target {
			return Pair{Start: start, End: end}
		} else if sum > target {
			end--
		} else {
			start++
		}
	}

	return Pair{-1, -1}
}
