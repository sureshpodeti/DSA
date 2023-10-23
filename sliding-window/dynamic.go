package sliding_window

import "math"

/*
//ar := []int{2, 1, 5, 2, 3, 2}
	//S := 7

	//ar := []int{2, 1, 5, 2, 8}
	//S := 7

	ar := []int{3, 4, 1, 1, 6}
	S := 8

	minLen := sliding_window.SmallestSubArrWithGivenSum(ar, S)

	fmt.Println(minLen)
*/

func SmallestSubArrWithGivenSum(ar []int, s int) int {
	minLength := len(ar)
	start := 0
	sum := 0

	for end := 0; end < len(ar); end++ {
		sum += ar[end]

		for sum >= s {
			minLength = int(math.Min(float64(minLength), float64(end-start+1)))
			sum -= ar[start]
			start++
		}
	}

	return minLength
}
