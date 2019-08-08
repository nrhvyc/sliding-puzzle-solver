package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	size, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("Could not parse puzzle size: %s", err)
	}

	puzzle := NewPuzzle(int(size))
	fmt.Println(puzzle.StartBoard.ToStringNotation())

	// visitedBoards := make(map[string]bool)
	pq := make(PriorityQueue, 1)
	firstPriority := puzzle.StartBoard.TotalScore(puzzle.GoalBoard.Tiles)
	node := &Node{
		value:    puzzle.StartBoard.ToStringNotation(),
		priority: firstPriority,
	}

	pq[0] = node
	heap.Init(&pq)

	threshold := firstPriority
	for threshold != 0 {
		fmt.Printf("Searching to threshold: %d...\n", threshold)
		threshold = search(pq, puzzle, threshold)
	}
}

/*
	Iterative Deepening A* Search
	Heuristics used include Hamming & Manhattan Distances
*/
func search(pq PriorityQueue, puzzle Puzzle, threshold int) int {
	/*
		minimum threshold will be calculated for use as the threshold
		in the next iteration
	*/
	minThreshold := 999999999999

	if puzzle.StartBoard.IsComplete(puzzle.GoalBoard) {
		puzzle.Print()
		puzzle.StartBoard.PrintWithGoal(puzzle.GoalBoard)
		fmt.Println("Goal!!!")
		return 0
	}

	for pq.Len() > 0 {
		time.Sleep(time.Second * 1 / 10)

		node := heap.Pop(&pq).(*Node)
		// fmt.Printf("%.2d:%s \n", node.priority, node.value)
		board := FromStringNotation(node.value)
		// fmt.Println("")
		// for pq.Len() > 0 {
		// 	node := heap.Pop(&pq).(*Node)
		// 	fmt.Printf("%.2d:%s \n", node.priority, node.value)
		// }
		// return 0

		// Debug output
		fmt.Println()
		// board.PrintWithGoal(puzzle.GoalBoard)
		// fmt.Printf("Node{priority:%.2d,value:%s}\n", node.priority, node.value)
		fmt.Printf("PriorityQueue size: %d  Path length: %d Path%v\n",
			pq.Len(), len(board.Path), board.Path)

		if board.IsComplete(puzzle.GoalBoard) {
			fmt.Println()
			board.PrintWithGoal(puzzle.GoalBoard)
			fmt.Printf("Node{priority:%.2d,value:%s}\n", node.priority, node.value)

			fmt.Printf("PriorityQueue size: %d  Path length: %d Path%v\n",
				pq.Len(), len(board.Path), board.Path)
			fmt.Println("Goal!!!")
			return 0
		}
		// var test []string
		for _, move := range board.PossibleMoves() {
			// Optimization to prevent queueing unnecessary nodes
			if board.PreviousMove == OppositeDirection(move.direction) && board.PreviousMove != "" {
				continue
			}

			// fmt.Printf("before move: %s OpenPosition: %v\n", move.direction, board.OpenPosition())
			newBoard := board.Move(move)
			totalScore := newBoard.TotalScore(puzzle.GoalBoard.Tiles)
			// fmt.Printf("after move: %s totalScore: %d OpenPosition: %v\n", move.direction, totalScore, newBoard.OpenPosition())
			// newBoard.Print()
			// Don't explore moves that are above the cost threshold
			if totalScore > threshold {
				if totalScore < minThreshold {
					minThreshold = totalScore
				}
				continue
			}

			heap.Push(&pq, &Node{
				value:    newBoard.ToStringNotation(),
				priority: totalScore,
			})
			// test = append(test, newBoard.ToStringNotation())
		}
		// return 0
		// if threshold > 2 {
		// 	fmt.Println("")
		// 	for pq.Len() > 0 {
		// 		node := heap.Pop(&pq).(*Node)
		// 		fmt.Printf("%.2d:%s \n", node.priority, node.value)
		// 		fmt.Println(test)
		// 	}
		// 	return 0
		// }
	}
	return minThreshold
}
