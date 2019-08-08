package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

//Board ...
type Board struct {
	Tiles        [][]int
	Path         []string
	PreviousMove string // Optimization to prevent queueing unnecessary nodes
}

//Position on a Board
type Position struct {
	x, y int
}

//Move for where the open position will be
type Move struct {
	direction string
}

//Remove an element from the slice
func remove(s []int, i int) []int {
	s[i] = s[0]
	return s[1:]
}

// OppositeDirection ...
func OppositeDirection(direction string) string {
	switch direction {
	case "L":
		return "R"
	case "R":
		return "L"
	case "U":
		return "D"
	case "D":
		return "U"
	}
	return ""
}

// DeepCopy - the underlying Tiles and Path need to be deep copied
// to prevent (b Board).Move() modifiying the original Board
func (b Board) DeepCopy() *Board {
	var newBoard Board

	// Deep copy Tiles
	n := len(b.Tiles)
	m := len(b.Tiles[0]) // in case the board is not square
	newBoard.Tiles = make([][]int, n)
	tile := make([]int, n*m)
	for i := range b.Tiles {
		start := i * m
		end := start + m
		newBoard.Tiles[i] = tile[start:end:end]
		copy(newBoard.Tiles[i], b.Tiles[i])
	}

	// Deep copy Path
	newBoard.Path = make([]string, len(b.Path))
	copy(newBoard.Path, b.Path)

	newBoard.PreviousMove = b.PreviousMove
	return &newBoard
}

// Move swaps the positon of tiles and returns a new Board
func (b Board) Move(m Move) Board {
	b = *b.DeepCopy() // Prevent modifying the original Board
	openPos := b.OpenPosition()

	b.PreviousMove = m.direction

	switch m.direction {
	case "L":
		b.Tiles[openPos.y][openPos.x] = b.Tiles[openPos.y][openPos.x-1]
		b.Tiles[openPos.y][openPos.x-1] = 0
	case "R":
		b.Tiles[openPos.y][openPos.x] = b.Tiles[openPos.y][openPos.x+1]
		b.Tiles[openPos.y][openPos.x+1] = 0
	case "U":
		b.Tiles[openPos.y][openPos.x] = b.Tiles[openPos.y-1][openPos.x]
		b.Tiles[openPos.y-1][openPos.x] = 0
	case "D":
		b.Tiles[openPos.y][openPos.x] = b.Tiles[openPos.y+1][openPos.x]
		b.Tiles[openPos.y+1][openPos.x] = 0
	}

	b.Path = append(b.Path, m.direction)

	return b
}

//NewRandomBoard ...
func NewRandomBoard(size int) Board {
	var board Board
	board.Tiles = make([][]int, size)
	for i := 0; i < size; i++ {
		board.Tiles[i] = make([]int, size)
	}

	// Create Seed Values
	var seedValues []int
	for j := 0; j < size*size; j++ {
		seedValues = append(seedValues, j)
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			randIndex := rand.Intn(len(seedValues))
			board.Tiles[x][y] = seedValues[randIndex]
			seedValues = remove(seedValues, randIndex)
		}
	}

	return board
}

// InversionCount - # of inversions for Tiles
// func (b Board) InversionCount() (count int) {
// 	for _, row := range b.Tiles {
// 		for i, tile := range row {
// 			for j, otherTile := range row {

// 			}
// 			if b.Tiles[y]
// 		}
// 	}
// 	return count
// }

// Print Board
func (b Board) Print() {
	for i := range b.Tiles {
		fmt.Print("|")
		fmt.Print(b.Row(i))
		fmt.Println("|")
	}
}

// PrintWithGoal ...
func (b Board) PrintWithGoal(goalBoard Board) {
	size := len(b.Tiles[0])
	for i := 0; i < size; i++ {
		line := "|"
		line += b.Row(i)

		if size/2 == i {
			line += "| -> |"
		} else {
			line += "|    |"
		}

		line += goalBoard.Row(i)
		line += "|"
		fmt.Println(line)
	}
}

//IsComplete ...
func (b Board) IsComplete(goalBoard Board) bool {
	for y, row := range b.Tiles {
		for x, val := range row {
			if val != goalBoard.Tiles[y][x] {
				return false
			}
		}
	}
	return true
}

//Row ...
func (b Board) Row(rowIndex int) string {
	var row string
	maxNum := len(b.Tiles[rowIndex]) * len(b.Tiles[rowIndex])

	for i, val := range b.Tiles[rowIndex] {
		if maxNum > 9 && val < 10 {
			row += " "
		}
		if val == 0 {
			row += "_"
		} else {
			row += strconv.Itoa(val)
		}
		if i < len(b.Tiles[rowIndex])-1 {
			row += " "
		}
	}
	return row
}

//OpenPosition that could be used for sliding
func (b Board) OpenPosition() (pos Position) {
	for y, row := range b.Tiles {
		for x, val := range row {
			if val == 0 {
				pos = Position{x: x, y: y}
				// b.Print()
				// fmt.Printf("OpenPosition: %d,%d\n", x, y)
				// time.Sleep(1 * time.Second)
				break
			}
		}
	}
	return
}

//PossibleMoves from the current board state
func (b Board) PossibleMoves() []Move {
	pos := b.OpenPosition()
	size := len(b.Tiles[0])

	var moves []Move
	// Left
	if pos.x > 0 {
		moves = append(moves, Move{
			direction: "L",
		})
	}
	// Right
	if pos.x <= size-2 {
		moves = append(moves, Move{
			direction: "R",
		})
	}
	// Up
	if pos.y > 0 {
		moves = append(moves, Move{
			direction: "U",
		})
	}
	// Down
	if pos.y <= size-2 {
		moves = append(moves, Move{
			direction: "D",
		})
	}
	fmt.Println(moves)
	return moves
}

//ToStringNotation returns the string notation for a Board
func (b Board) ToStringNotation() string {
	var stringNotation string
	size := len(b.Tiles[0])
	for y, row := range b.Tiles {
		for x, val := range row {
			stringNotation += strconv.Itoa(val)
			if x != size-1 {
				stringNotation += ","
			}
		}
		if y != size-1 {
			stringNotation += "/"
		}
	}

	stringNotation += "#"

	// Path Notation - if applicable
	for _, direction := range b.Path {
		stringNotation += fmt.Sprintf("%s",
			direction)
		// if i < len(b.Path)-1 {
		// 	stringNotation += "|"
		// }
	}

	stringNotation += "#"

	// PreviousOpenTile
	// stringNotation += fmt.Sprintf("%d,%d",
	// 	b.PreviousOpenTile.x,
	// 	b.PreviousOpenTile.y)
	stringNotation += b.PreviousMove

	return stringNotation
}

// FromStringNotation returns a Board from the string notation
func FromStringNotation(notation string) Board {
	var board Board
	notations := strings.Split(notation, "#")

	// Extract Board.Tiles
	rows := strings.Split(notations[0], "/")
	board.Tiles = make([][]int, len(rows))

	for y, row := range rows {
		vals := strings.Split(row, ",")
		board.Tiles[y] = make([]int, len(vals))

		for x, val := range vals {
			val, err := strconv.ParseInt(val, 10, 64)
			board.Tiles[y][x] = int(val)
			if err != nil {
				panic(err)
			}
		}
	}

	// parse string coordinate into int
	// parse := func(coord string) int {
	// 	val, err := strconv.ParseInt(coord, 10, 64)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	return int(val)
	// }

	// Extract Path - if applicable
	if notations[1] != "" {
		board.Path = strings.Split(notations[1], "")
		// moves := strings.Split(notations[1], "|")
		// for _, move := range moves {
		// 	coords := strings.Split(move, ",")

		// 	board.Path = append(board.Path, Move{
		// 		from: Position{
		// 			x: parse(coords[0]),
		// 			y: parse(coords[1])},
		// 		to: Position{
		// 			x: parse(coords[2]),
		// 			y: parse(coords[3])},
		// 	})
		// }
	}

	// Extract PreviousMove - if applicable
	board.PreviousMove = notations[2]

	return board
}

// ManhattanDistance for each tile
func (b Board) ManhattanDistance(goal [][]int) int {
	var distance float64
	distance = 0
	for y, row := range b.Tiles {
		for x, tile := range row {
			// Don't consider the empty space this would
			// make the heuristic not be an underestimate
			if tile == 0 {
				continue
			}
			goalTilePos := b.PositionForTile(tile)
			distance += math.Abs(float64(y-goalTilePos.x)) + math.Abs(float64(x-goalTilePos.y))
		}
	}
	return int(distance)
}

// PositionForTile - Search for position of tile value
func (b Board) PositionForTile(tile int) (pos Position) {
	for y, row := range b.Tiles {
		for x, tileVal := range row {
			if tile == tileVal {
				pos = Position{x: x, y: y}
				return
			}
		}
	}
	return pos
}

// HammingDistance heuristic
func (b Board) HammingDistance(goal [][]int) (misplacedCount int) {
	for y, row := range b.Tiles {
		for x, val := range row {
			if val == 0 {
				continue
			}
			if val != goal[y][x] {
				misplacedCount++
			}
		}
	}
	return
}

// HeuristicScore is the overall he
func (b Board) HeuristicScore(goal [][]int) int {
	// return b.ManhattanDistance(goal) + b.HammingDistance(goal)
	return b.HammingDistance(goal)
	// return 1
}

// TotalScore - lower values are more important
// f(n) = g(n) + h(n)
func (b Board) TotalScore(goal [][]int) int {
	h := b.HeuristicScore(goal)
	// fmt.Printf("%d + %d\n", pathLength, h)
	// time.Sleep(time.Second * 1)
	return len(b.Path) + h
}
