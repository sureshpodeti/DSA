package kWayMerge

/*
	minHeap := &kWayMerge.MinHeap{
			kWayMerge.Node{Data: 2, Index: 0},
			kWayMerge.Node{Data: 3, Index: 1},
			kWayMerge.Node{Data: 1, Index: 2},
		}
		heap.Init(minHeap)

		for minHeap.Len() > 0 {
			fmt.Println(heap.Pop(minHeap))
		}
*/
type Node struct {
	Data  int
	Index int
}

type MinHeap []Node

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].Data < m[j].Data
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MinHeap) Push(x any) {
	m = append(m, x.(Node))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}
