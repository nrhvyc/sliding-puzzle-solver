package main

import (
	"container/heap"
)

// An Node is something we manage in a priority queue.
type Node struct {
	value    string // The value of the node; arbitrary.
	priority int    // The priority of the node in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the node in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, cost priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

//Swap ...
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

//Push ...
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

//Pop ...
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

//Update modifies the priority and value of an Node in the queue.
func (pq *PriorityQueue) Update(node *Node, value string, priority int) {
	node.value = value
	node.priority = priority
	heap.Fix(pq, node.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an node,
// and then removes the items in priority order.
// func main() {
// 	// Some items and their priorities.
// 	items := map[string]int{
// 		"banana": 3, "apple": 2, "pear": 4,
// 	}

// 	// Create a priority queue, put the items in it, and
// 	// establish the priority queue (heap) invariants.
// 	pq := make(PriorityQueue, len(items))
// 	i := 0
// 	for value, priority := range items {
// 		pq[i] = &Node{
// 			value:    value,
// 			priority: priority,
// 			index:    i,
// 		}
// 		i++
// 	}
// 	heap.Init(&pq)

// 	// Insert a new node and then modify its priority.
// 	node := &Node{
// 		value:    "orange",
// 		priority: 1,
// 	}
// 	heap.Push(&pq, node)
// 	pq.Update(node, node.value, 5)

// 	// Take the items out; they arrive in decreasing priority order.
// 	for pq.Len() > 0 {
// 		node := heap.Pop(&pq).(*Node)
// 		fmt.Printf("%.2d:%s ", node.priority, node.value)
// 	}
// }
