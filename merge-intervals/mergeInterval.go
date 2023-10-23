package merge_intervals

import (
	"math"
	"sort"
)

/*
	//intervals := [][]int{
	//	{7, 9},
	//	{1, 4},
	//	{2, 5},
	//}

	//intervals := [][]int{
	//	{6, 7},
	//	{2, 4},
	//	{5, 9},
	//}

	intervals := [][]int{
		{1, 4},
		{2, 6},
		{3, 5},
	}

	mergedIntervals := merge_intervals.Merge(intervals)
	fmt.Println(mergedIntervals)
*/

func Merge(intervals [][]int) [][]int {
	//first sort the array
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var mergedIntervals [][]int

	prevIntervalStart, prevIntervalEnd := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= prevIntervalEnd {
			// Overlapping windows
			prevIntervalEnd = int(math.Max(float64(intervals[i][1]), float64(prevIntervalEnd)))
		} else {
			// Non-overlapping windows
			mergedIntervals = append(mergedIntervals, []int{prevIntervalStart, prevIntervalEnd})
			prevIntervalStart, prevIntervalEnd = intervals[i][0], intervals[i][1]
		}
	}
	mergedIntervals = append(mergedIntervals, []int{prevIntervalStart, prevIntervalEnd})

	return mergedIntervals

}
