package main

import (
	"reflect"
	"testing"
)

func TestLoadSet(t *testing.T) {
	tests := []struct {
		input string
		want  Set
	}{
		{
			input: "1 green, 3 red, 6 blue",
			want:  Set{Green: 1, Red: 3, Blue: 6},
		},
		{
			input: "6 red, 1 blue, 3 green",
			want:  Set{Green: 3, Red: 6, Blue: 1},
		},
		{
			input: "1 green, 1 blue",
			want:  Set{Green: 1, Red: 0, Blue: 1},
		},
	}
	for _, test := range tests {
		got := Set{}
		got.Load(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("import set(%+v) = %+v, but want %+v", test.input, got, test.want)
		}
	}
}

func TestLoadGame(t *testing.T) {
	tests := []struct {
		input string
		want  Game
	}{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",

			want: Game{ID: 1, Sets: []Set{{Green: 0, Red: 4, Blue: 3}, {Green: 2, Red: 1, Blue: 6}, {Green: 2, Red: 0, Blue: 0}}},
		},
		{
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",

			want: Game{ID: 2, Sets: []Set{{Green: 2, Red: 0, Blue: 1}, {Green: 3, Red: 1, Blue: 4}, {Green: 1, Red: 0, Blue: 1}}},
		},
	}
	for _, test := range tests {
		got := Game{}
		got.Load(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("import game(%+v) = \ngot:  %+v\nwant: %+v", test.input, got, test.want)
		}
	}
}

func TestGameIsPossible(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:  true,
		},
		{
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:  true,
		},
		{
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:  false,
		},
		{
			input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want:  false,
		},
		{
			input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want:  true,
		},
	}
	for _, test := range tests {
		got := Game{}
		got.Load(test.input)
		if got.IsPossible(12, 13, 14) != test.want {
			t.Errorf("is possible game(%+v) = \ngot:  %+v\nwant: %+v", test.input, got, test.want)
		}
	}
}
