package main

import (
	"container/heap"
)

// Node holds a stringified Board state and it's corresponding priority
type Node struct {
	value    string // The value of the node; arbitrary.
	priority int    // The priority of the node in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the node in the heap.
}

// PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest,
	// cost priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

// Swap ...
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push ...
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

// Pop ...
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

// Update modifies the priority and value of an Node in the queue.
func (pq *PriorityQueue) Update(node *Node, value string, priority int) {
	node.value = value
	node.priority = priority
	heap.Fix(pq, node.index)
}
