package main

import "fmt"

//Puzzle ...
type Puzzle struct {
	StartBoard Board // starts as the initial state
	GoalBoard  Board
	Size       int
}

//Print Puzzle
func (p Puzzle) Print() {
	for i := 0; i < p.Size; i++ {
		line := "|"
		line += p.StartBoard.Row(i)

		if p.Size/2 == i {
			line += "| -> |"
		} else {
			line += "|    |"
		}

		line += p.GoalBoard.Row(i)
		line += "|"
		fmt.Println(line)

	}
}

// IsSolvable Puzzle
/*
1. If the grid width is odd,
	then the number of inversions for solvable is even.
2. If the grid width is even,
	and the blank is on an even row counting from the bottom
	(second-last, fourth-last etc)
	then the number of inversions for solvable is odd.
3. If the grid width is even,
	and the blank is on an odd row counting from the bottom
	(last, third-last, fifth-last etc)
	then the number of inversions for solvable is even.
*/
func (p Puzzle) IsSolvable() bool {
	startInversions := p.StartBoard.InversionCount()
	size := len(p.StartBoard.Tiles[0])
	if size%2 == 1 && startInversions%2 == 0 {
		return true
	}
	if size%2 == 0 {
		openPos := p.StartBoard.OpenPosition()
		y := openPos.y + 1
		fmt.Println(openPos)
		fmt.Println(startInversions)

		if (size-y)%2 == 1 && startInversions%2 == 1 {
			return true
		}
		if (size-y)%2 == 0 && startInversions%2 == 0 {
			return true
		}
	}
	return false
}

//NewPuzzle ...
func NewPuzzle(size int) Puzzle {
	// puzzle := Puzzle{
	// 	StartBoard: NewRandomBoard(size),
	// 	GoalBoard:  NewRandomBoard(size),
	// 	Size:       size,
	// }

	// Optimal move count 30
	// ULLURDDRUULDDLUURDDRUULDDLURRD or ULULDDRUULDDRUURDDLUURDLULDRDR
	puzzle := Puzzle{
		StartBoard: Board{Tiles: [][]int{
			{8, 7, 6},
			{5, 4, 3},
			{2, 1, 0}},
			PreviousMove: ""},
		GoalBoard: Board{Tiles: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 0}}},
		Size: size,
	}

	// Optimal move count 25
	// DLULURDRULDDLUURDRDLLURRD
	// puzzle := Puzzle{
	// 	StartBoard: Board{Tiles: [][]int{
	// 		{8, 7, 4},
	// 		{3, 2, 0},
	// 		{6, 5, 1}},
	// 		PreviousMove: ""},
	// 	GoalBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3},
	// 		{4, 5, 6},
	// 		{7, 8, 0}}},
	// 	Size: size,
	// }
	// 4  1  3
	// 0  2  6
	// 7  5  8
	// 	0  1  2  3
	// 	5  6  7  4
	// 	9 10 11  8
	//    13 14 15 12

	// puzzle := Puzzle{
	// 	StartBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3, 4, 5},
	// 		{12, 6, 8, 9, 10},
	// 		{0, 7, 13, 19, 14},
	// 		{11, 16, 17, 18, 15},
	// 		{21, 22, 23, 24, 20}},
	// 		PreviousMove: ""},
	// 	GoalBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3, 4, 5},
	// 		{6, 7, 8, 9, 10},
	// 		{11, 12, 13, 14, 15},
	// 		{16, 17, 18, 19, 20},
	// 		{21, 22, 23, 24, 25}}},
	// 	Size: size,
	// }

	// 15 moves
	// puzzle := Puzzle{
	// 	StartBoard: Board{Tiles: [][]int{
	// 		{5, 1, 7, 3},
	// 		{9, 2, 11, 4},
	// 		{13, 6, 15, 8},
	// 		{0, 10, 14, 12}},
	// 		PreviousMove: ""},
	// 	GoalBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3, 4},
	// 		{5, 6, 7, 8},
	// 		{9, 10, 11, 12},
	// 		{13, 14, 15, 0}}},
	// 	Size: size,
	// }

	// 26 inversions, solvable
	// inversions: 4 + 0 + 4 + 1  +  4 + 0 + 4 + 0  + 4 + 0 + 4 + 0  +  0 + 0 + 1 + 0
	// puzzle := Puzzle{
	// 	StartBoard: Board{Tiles: [][]int{
	// 		{5, 1, 7, 3},
	// 		{9, 2, 11, 4},
	// 		{13, 6, 15, 8},
	// 		{0, 10, 14, 12}},
	// 		PreviousMove: ""},
	// 	GoalBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3, 4},
	// 		{5, 6, 7, 8},
	// 		{9, 10, 11, 12},
	// 		{13, 14, 15, 0}}},
	// 	Size: size,
	// }

	// Optimal move count 9		LURDLLDRR
	// puzzle := Puzzle{
	// 	StartBoard: Board{Tiles: [][]int{
	// 		{1, 6, 2},
	// 		{5, 3, 0},
	// 		{4, 7, 8}},
	// 		PreviousMove: ""},
	// 	GoalBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3},
	// 		{4, 5, 6},
	// 		{7, 8, 0}}},
	// 	Size: size,
	// }

	// Optimal move count 5		DLURD
	// puzzle := Puzzle{
	// 	StartBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3},
	// 		{4, 8, 0},
	// 		{7, 6, 5}}},
	// 	GoalBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3},
	// 		{4, 5, 6},
	// 		{7, 8, 0}}},
	// 	Size: size,
	// }

	// Optimal move count 4		DLURD
	// puzzle := Puzzle{
	// 	StartBoard: Board{
	// 		Tiles: [][]int{
	// 			{1, 2, 3},
	// 			{7, 4, 5},
	// 			{0, 8, 6}},
	// 		PreviousOpenTile: Position{x: 0, y: 2}},
	// 	GoalBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3},
	// 		{4, 5, 6},
	// 		{7, 8, 0}}},
	// 	Size: size,
	// }

	// Optimal move count 4		DLURD
	// puzzle := Puzzle{
	// 	StartBoard: Board{
	// 		Tiles: [][]int{
	// 			{1, 2, 3},
	// 			{4, 0, 5},
	// 			{7, 8, 6}},
	// 		PreviousOpenTile: Position{x: 0, y: 0}},
	// 	GoalBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3},
	// 		{4, 5, 6},
	// 		{7, 8, 0}}},
	// 	Size: size,
	// }

	// puzzle := Puzzle{
	// 	StartBoard: Board{Tiles: [][]int{
	// 		{1, 2, 3},
	// 		{8, 0, 4},
	// 		{7, 6, 5}}},
	// 	GoalBoard: Board{Tiles: [][]int{
	// 		{2, 8, 1},
	// 		{0, 4, 3},
	// 		{7, 6, 5}}},
	// 	Size: size,
	// }

	// puzzle := Puzzle{
	// 	StartBoard: Board{Tiles: [][]int{
	// 		{0, 1},
	// 		{2, 3}}},
	// 	GoalBoard: Board{Tiles: [][]int{
	// 		{0, 2},
	// 		{1, 3}}},
	// 	Size: size,
	// }

	if puzzle.IsSolvable() == false {
		panic("Puzzle provided is not solvable")
	}

	return puzzle
}
