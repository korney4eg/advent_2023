package main

import (
	"testing"
)

func TestGetAdjacentNumbers(t *testing.T) {
	// Test that the function returns the correct adjacent numbers
	// for a given number
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: 4361,
		},
		{
			input: `467..114..
...*......`,
			want: 467,
		},
		{
			input: `....90*12...`,
			want:  102,
		},
		{
			input: `2.2......12.
.*.........*
1.1.......56`,
			want: 74,
		},
	}
	for _, test := range tests {
		got, _ := getAdjacentNumbers(test.input)
		if got != test.want {
			t.Errorf("got %v, but want %v", got, test.want)
		}
	}
}
func TestGetProperGears(t *testing.T) {
	// Test that the function returns the correct adjacent numbers
	// for a given number
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: 467835,
		},
	}
	for _, test := range tests {
		got, _ := getGears(test.input)
		if got != test.want {
			t.Errorf("got %v, but want %v", got, test.want)
		}
	}
}
