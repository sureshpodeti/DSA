package topologicalSort

import "container/list"

/*
// 4
	//edges := [][]int{
	//	{3, 2},
	//	{3, 0},
	//	{2, 0},
	//	{2, 1},
	//}

	// 5
	//edges := [][]int{
	//	{4, 2},
	//	{4, 3},
	//	{2, 0},
	//	{2, 1},
	//	{3, 1},
	//}

	// 7
	edges := [][]int{
		{6, 4},
		{6, 2},
		{5, 3},
		{5, 4},
		{3, 0},
		{3, 1},
		{3, 2},
		{4, 1},
	}

	sortedOrder := topologicalSort.TopologicalSort(edges)
	fmt.Println(sortedOrder)
*/

const nVertices = 7

func TopologicalSort(edges [][]int) []int {
	var sortedOrder []int

	var inDegree [nVertices]int

	//Calculate adjacency list
	var graph map[int][]int

	queue := list.New()

	graph = make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		vertex := edges[i][0]
		reachable := edges[i][1]
		graph[vertex] = append(graph[vertex], reachable)
		inDegree[reachable]++
	}

	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			//Vertex with no in-degree
			queue.PushBack(i)
		}
	}

	e := queue.Front()

	for e != nil {
		// Collect this element in the sortedOrder
		sortedOrder = append(sortedOrder, e.Value.(int))

		// update it's children inDegree values
		children := graph[e.Value.(int)]

		for _, c := range children {
			inDegree[c]--

			if inDegree[c] == 0 {
				queue.PushBack(c)
			}
		}
		queue.Remove(e)
		e = queue.Front()
	}

	return sortedOrder

}
