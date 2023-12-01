package main

import (
	"testing"
)

func TestCalibrate(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	want := 142
	if got := Calibrate(input, false); got != want {
		t.Errorf("calibrate (%q) = %d, but want %d", input, got, want)
	}
	input = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	want = 281
	if got := Calibrate(input, true); got != want {
		t.Errorf("calibrate (%q) = %d, but want %d", input, got, want)
	}

}

func TestCalibrateLine(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "a1dfasd4",
			want:  14,
		},
		{
			input: "4daffa6d",
			want:  46,
		},
	}
	for _, test := range tests {
		if got := CalibrateLine(test.input, false); got != test.want {
			t.Errorf("calibrate line(%q) = %d, but want %d", test.input, got, test.want)
		}
	}
}
func TestCalibrateLineWordNums(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "eighthree",
			want:  83,
		},
		{
			input: "sevenine",
			want:  79,
		},
		{
			input: "zoneight234",
			want:  14,
		},
		{
			input: "eightwothree",
			want:  83,
		},
		{
			input: "xtwone3four",
			want:  24,
		},
		{
			input: "4nineeightseven2",
			want:  42,
		},
	}
	for _, test := range tests {
		if got := CalibrateLine(test.input, true); got != test.want {
			t.Errorf("calibrate line(%q) = %d, but want %d", test.input, got, test.want)
		}
	}
}
