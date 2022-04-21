package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 3, 4, 7, 9}, 9},
	}

	for _, test := range tests {
		if got := max(test.input...); got != test.want {
			t.Errorf("max(%v) = %d, want %d\n", test.input, got, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 3, 4, 7, 9}, 1},
	}

	for _, test := range tests {
		if got := min(test.input...); got != test.want {
			t.Errorf("min(%v) = %d, want %d\n", test.input, got, test.want)
		}
	}
}

func TestAtLeastOneParameterMax2(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 3}, 3},
	}

	for _, test := range tests {
		if got := max2(test.input[0], test.input[1:]...); got != test.want {
			t.Errorf("max2(%v) = %d, want %d\n", test.input, got, test.want)
		}
	}
}

func TestAtLeastOneParameterMin2(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 3}, 1},
	}
	for _, test := range tests {
		if got := min2(test.input[0], test.input[1:]...); got != test.want {
			t.Errorf("min2(%v) = %d, want %d\n", test.input, got, test.want)
		}
	}
}

func TestZeroParameterMax(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{}, 0},
	}

	for _, test := range tests {
		if got := max(test.input...); got != test.want {
			t.Errorf("max(%v) = %d, want %d\n", test.input, got, test.want)
		}
	}
}

func TestZeroParameterMin(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{}, 0},
	}

	for _, test := range tests {
		if got := min(test.input...); got != test.want {
			t.Errorf("min(%v) = %d, want %d\n", test.input, got, test.want)
		}
	}
}
