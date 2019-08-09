package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	size, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("Could not parse puzzle size: %s", err)
	}

	puzzle := NewPuzzle(int(size))
	fmt.Println(puzzle.StartBoard.ToStringNotation())

	firstPriority := puzzle.StartBoard.TotalScore(puzzle.GoalBoard.Tiles)
	node := &Node{
		value:    puzzle.StartBoard.ToStringNotation(),
		priority: firstPriority,
	}

	threshold := firstPriority
	for threshold != 0 {
		pq := make(PriorityQueue, 1)
		pq[0] = node
		heap.Init(&pq)
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
		// time.Sleep(time.Second * 1 / 10)

		node := heap.Pop(&pq).(*Node)
		board := FromStringNotation(node.value)

		// Debug output
		// fmt.Println()
		// fmt.Printf("PriorityQueue size: %d  Path length: %d Path%v\n",
		// 	pq.Len(), len(board.Path), board.Path)

		if board.IsComplete(puzzle.GoalBoard) {
			fmt.Println()
			board.PrintWithGoal(puzzle.GoalBoard)
			fmt.Printf("Node{priority:%.2d,value:%s}\n", node.priority, node.value)

			fmt.Printf("PriorityQueue size: %d  Path length: %d Path%v\n",
				pq.Len(), len(board.Path), board.Path)
			fmt.Println("Goal!!!")
			return 0
		}

		for _, move := range board.PossibleMoves() {
			// Optimization to prevent queueing unnecessary nodes
			if board.PreviousMove == OppositeDirection(move.direction) && board.PreviousMove != "" {
				continue
			}

			newBoard := board.Move(move)
			totalScore := newBoard.TotalScore(puzzle.GoalBoard.Tiles)

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
		}
	}
	return minThreshold
}
