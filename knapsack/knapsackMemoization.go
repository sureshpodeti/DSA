package knapsack

import "math"

/*
weights := []int{1, 2, 3, 5}
	profits := []int{1, 6, 10, 16}

	//capacity := 7

	capacity := 6

	answer := knapsack.Knapsack(weights, profits, capacity)

	fmt.Printf("Answer: %d\n", answer)

*/

func KnapsackMemo(weights, profits []int, capacity int) int {
	return knapsackMemoUtil(weights, profits, capacity, 0)
}

func knapsackMemoUtil(weights, profits []int, capacity, index int) int {
	if capacity <= 0 || index >= len(weights) {
		return 0
	}
	profit1 := 0
	if capacity >= weights[index] {
		profit1 = profits[index] + knapsackMemoUtil(weights, profits, capacity-weights[index], index+1)
	}
	profit2 := knapsackMemoUtil(weights, profits, capacity, index+1)

	return int(math.Max(float64(profit1), float64(profit2)))

}
