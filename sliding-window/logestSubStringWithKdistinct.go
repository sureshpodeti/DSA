package sliding_window

import (
	"fmt"
	"math"
)

/*
	//str := "araaci"
	//K := 2

	//str := "araaci"
	//K := 1

	str := "cbbebi"
	K := 3

	result := sliding_window.LongestSubStrWithKDistictChars(str, K)
	fmt.Println(result)
*/

func LongestSubStrWithKDistictChars(str string, k int) int {

	mp := make(map[string]int)

	start := 0

	result := 1

	for end := 0; end < len(str); end++ {
		mp[string(str[end])]++

		if len(mp) > k {
			result = int(math.Max(float64(result), float64(end-start)))

			if mp[string(str[start])] == 1 {
				delete(mp, string(str[start]))
			} else {
				mp[string(str[start])]--
			}

			start++
		}

	}

	fmt.Println(mp)

	return result
}
