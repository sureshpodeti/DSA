package knapsack

import "math"

/*
weights := []int{1, 2, 3, 5}
	profits := []int{1, 6, 10, 16}
	//capacity := 6
	capacity := 7
	answer := knapsack.KnapsackTabular(weights, profits, capacity)

	fmt.Printf("Answer: %d\n", answer)
*/

func KnapsackTabular(weights, profits []int, capacity int) int {
	nWeights := len(weights)
	//nProfits := len(profits)

	dp := make([][]int, nWeights)

	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	//initialize capacity = 0; dp[i][0] = 0
	for i := 0; i < nWeights; i++ {
		dp[i][0] = 0
	}

	for c := 0; c <= capacity; c++ {
		if weights[0] <= c {
			dp[0][c] = profits[0]
		}
	}

	for i := 1; i < nWeights; i++ {
		for j := 1; j <= capacity; j++ {
			profit1, profit2 := 0, 0
			if weights[i] <= j {
				profit1 = profits[i] + dp[i-1][j-weights[i]]
			}

			profit2 = dp[i-1][j]

			dp[i][j] = int(math.Max(float64(profit1), float64(profit2)))

		}
	}
	return dp[nWeights-1][capacity]
}
