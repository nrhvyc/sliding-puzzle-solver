package main

import (
	"reflect"
	"testing"
)

func Test_remove(t *testing.T) {
	type args struct {
		s []int
		i int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Move(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	type args struct {
		m Move
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Board
	}{
		{
			name: "center move right",
			fields: fields{
				Tiles: [][]int{
					{1, 2, 3},
					{4, 0, 5},
					{7, 8, 6}},
			},
			args: args{m: Move{direction: "R"}},
			want: Board{
				Tiles: [][]int{
					{1, 2, 3},
					{4, 5, 0},
					{7, 8, 6}},
				Path:         []string{"R"},
				PreviousMove: "R",
			},
		}, {
			name: "center move down",
			fields: fields{
				Tiles: [][]int{
					{1, 2, 3},
					{4, 0, 5},
					{7, 8, 6}},
			},
			args: args{m: Move{direction: "D"}},
			want: Board{
				Tiles: [][]int{
					{1, 2, 3},
					{4, 8, 5},
					{7, 0, 6}},
				Path:         []string{"D"},
				PreviousMove: "D",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			if got := b.Move(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.Move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRandomBoard(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want Board
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRandomBoard(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandomBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Print(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			b.Print()
		})
	}
}

func TestBoard_PrintWithGoal(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	type args struct {
		goalBoard Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			b.PrintWithGoal(tt.args.goalBoard)
		})
	}
}

func TestBoard_IsComplete(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	type args struct {
		goalBoard Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			if got := b.IsComplete(tt.args.goalBoard); got != tt.want {
				t.Errorf("Board.IsComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Row(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			if got := b.Row(tt.args.rowIndex); got != tt.want {
				t.Errorf("Board.Row() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_OpenPosition(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	tests := []struct {
		name    string
		fields  fields
		wantPos Position
	}{
		{
			name: "misplaced 1",
			fields: fields{
				Tiles: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 0, 8}},
			},
			wantPos: Position{x: 1, y: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			if gotPos := b.OpenPosition(); !reflect.DeepEqual(gotPos, tt.wantPos) {
				t.Errorf("Board.OpenPosition() = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}

func TestBoard_PossibleMoves(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	tests := []struct {
		name   string
		fields fields
		want   []Move
	}{
		{
			name: "center",
			fields: fields{
				Tiles: [][]int{
					{1, 2, 3},
					{4, 0, 6},
					{7, 8, 5}},
			},
			want: []Move{
				Move{direction: "L"},
				Move{direction: "R"},
				Move{direction: "U"},
				Move{direction: "D"},
			},
		}, {
			name: "bottom right",
			fields: fields{
				Tiles: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 0}},
			},
			want: []Move{
				Move{direction: "L"},
				Move{direction: "U"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles: tt.fields.Tiles,
			}
			if got := b.PossibleMoves(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.PossibleMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_ToStringNotation(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			if got := b.ToStringNotation(); got != tt.want {
				t.Errorf("Board.ToStringNotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromStringNotation(t *testing.T) {
	type args struct {
		notation string
	}
	tests := []struct {
		name string
		args args
		want Board
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromStringNotation(tt.args.notation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromStringNotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_ManhattanDistance(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	type args struct {
		goal [][]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "misplaced 0",
			fields: fields{
				Tiles: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 0}},
			},
			args: args{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			if got := b.ManhattanDistance(tt.args.goal); got != tt.want {
				t.Errorf("Board.ManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_PositionForTile(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	type args struct {
		tile int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantPos Position
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			if gotPos := b.PositionForTile(tt.args.tile); !reflect.DeepEqual(gotPos, tt.wantPos) {
				t.Errorf("Board.PositionForTile() = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}

func TestBoard_HammingDistance(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	type args struct {
		goal [][]int
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantMisplacedCount int
	}{
		{
			name: "misplaced 0",
			fields: fields{
				Tiles: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 0}},
			},
			args: args{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0}}},
			wantMisplacedCount: 0,
		}, {
			name: "misplaced 1",
			fields: fields{
				Tiles: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 0, 8}},
			},
			args: args{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0}}},
			wantMisplacedCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles: tt.fields.Tiles,
			}
			if gotMisplacedCount := b.HammingDistance(tt.args.goal); gotMisplacedCount != tt.wantMisplacedCount {
				t.Errorf("Board.HammingDistance() = %v, want %v", gotMisplacedCount, tt.wantMisplacedCount)
			}
		})
	}
}

func TestBoard_HeuristicScore(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	type args struct {
		goal [][]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			if got := b.HeuristicScore(tt.args.goal); got != tt.want {
				t.Errorf("Board.HeuristicScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_TotalScore(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	type args struct {
		goal [][]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles:        tt.fields.Tiles,
				Path:         tt.fields.Path,
				PreviousMove: tt.fields.PreviousMove,
			}
			if got := b.TotalScore(tt.args.goal); got != tt.want {
				t.Errorf("Board.TotalScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
