package main

import "testing"

func TestPuzzle_IsSolvable(t *testing.T) {
	type fields struct {
		StartBoard Board
		GoalBoard  Board
		Size       int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "49 inversions, solvable",
			fields: fields{
				StartBoard: Board{Tiles: [][]int{
					{12, 1, 10, 2},
					{7, 11, 4, 14},
					{5, 0, 9, 15},
					{8, 13, 6, 3}},
				},
				GoalBoard: Board{Tiles: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 0}},
				},
				Size: 4,
			},
			want: true,
		}, {
			// inversions:
			// 4 + 0 + 4 + 1 +
			// 4 + 0 + 4 + 0 +
			// 4 + 0 + 4 + 0 +
			// 0 + 0 + 1 + 0
			name: "26 inversions, solvable",
			fields: fields{
				StartBoard: Board{Tiles: [][]int{
					{5, 1, 7, 3},
					{9, 2, 11, 4},
					{13, 6, 15, 8},
					{0, 10, 14, 12}},
				},
				GoalBoard: Board{Tiles: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 0}},
				},
				Size: 4,
			},
			want: true,
		}, {
			name: "9-puzzle, solvable",
			fields: fields{
				StartBoard: Board{Tiles: [][]int{
					{1, 6, 2},
					{5, 3, 0},
					{4, 7, 8}},
				},
				GoalBoard: Board{Tiles: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 0}},
				},
				Size: 4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Puzzle{
				StartBoard: tt.fields.StartBoard,
				GoalBoard:  tt.fields.GoalBoard,
				Size:       tt.fields.Size,
			}
			if got := p.IsSolvable(); got != tt.want {
				t.Errorf("Puzzle.IsSolvable() = %v, want %v", got, tt.want)
			}
		})
	}
}
