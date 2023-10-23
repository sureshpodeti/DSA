package sliding_window

import "math"

/*
	//b := []byte{'A', 'B', 'C', 'A', 'C'}

	b := []byte{'A', 'B', 'C', 'B', 'B', 'C'}

	result := sliding_window.FruitBasket(b)

	fmt.Println(result)
*/

func FruitBasket(str []byte) int {
	mp := make(map[byte]int)

	start := 0
	result := 0

	for end := 0; end < len(str); end++ {
		mp[str[end]]++

		for len(mp) > 2 {

			if mp[str[start]] > 1 {
				mp[str[start]]--
			} else {
				delete(mp, str[start])
			}

			start++
		}
		result = int(math.Max(float64(result), float64(end-start+1)))
	}

	return result
}
