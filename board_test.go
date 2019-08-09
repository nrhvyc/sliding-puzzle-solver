package main

import (
	"reflect"
	"testing"
)

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

func TestBoard_InversionCount(t *testing.T) {
	type fields struct {
		Tiles        [][]int
		Path         []string
		PreviousMove string
	}
	tests := []struct {
		name      string
		fields    fields
		wantCount int
	}{
		{
			name: "inversion count 49",
			fields: fields{
				Tiles: [][]int{
					{12, 1, 10, 2},
					{7, 11, 4, 14},
					{5, 0, 9, 15},
					{8, 13, 6, 3}},
			},
			wantCount: 49,
		}, {
			// inversions:
			// 4 + 0 + 4 + 1 +
			// 4 + 0 + 4 + 0 +
			// 4 + 0 + 4 + 0 +
			// 0 + 0 + 1 + 0
			name: "inversion count 26",
			fields: fields{
				Tiles: [][]int{
					{5, 1, 7, 3},
					{9, 2, 11, 4},
					{13, 6, 15, 8},
					{0, 10, 14, 12}},
			},
			wantCount: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Tiles: tt.fields.Tiles,
			}
			if gotCount := b.InversionCount(); gotCount != tt.wantCount {
				t.Errorf("Board.InversionCount() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
