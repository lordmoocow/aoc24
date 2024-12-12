package utils

import (
	"fmt"
	"testing"
)

func TestGridGet(t *testing.T) {
	grid := Grid[int]{
		cells: []int{7, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 9, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 8},
		width: 5,
	}
	tests := []struct {
		input  Point
		result int
		ok     bool
	}{
		{input: Point{0, 0}, ok: true, result: 7},
		{input: Point{4, 4}, ok: true, result: 8},
		{input: Point{-1, 2}, ok: false, result: 0},
		{input: Point{2, -1}, ok: false, result: 0},
		{input: Point{-1, -1}, ok: false, result: 0},
		{input: Point{3, 2}, ok: true, result: 9},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			t.Parallel()

			if got, ok := grid.Get(test.input); ok != test.ok || got != test.result {
				t.Fatalf("Grid.Get(%q) returned %q, %t; expected %q, %t", test.input, got, ok, test.result, test.ok)
			}
		})
	}
}

func TestGridSet(t *testing.T) {
	grid := Grid[int]{
		cells: []int{7, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 9, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 8},
		width: 5,
	}
	tests := []struct {
		input Point
		value int
		index int
		ok    bool
	}{
		{input: Point{0, 0}, ok: true, index: 0, value: 100},
		{input: Point{4, 4}, ok: true, index: 24, value: 101},
		{input: Point{-1, 2}, ok: false, index: 0, value: 0},
		{input: Point{2, -1}, ok: false, index: 0, value: 0},
		{input: Point{-1, -1}, ok: false, index: 0, value: 0},
		{input: Point{3, 2}, ok: true, index: 13, value: 102},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			t.Parallel()
			ok := grid.Set(test.input, test.value)
			if ok != test.ok {
				t.Fatalf("Grid.Set(%q) returned %t; expected %t. ", test.input, ok, test.ok)
			}
			if ok && grid.cells[test.index] != test.value {
				t.Fatalf("Grid.Set(%q) returned %t but got value %q; expected %q. ", test.input, ok, grid.cells[test.index], test.value)
			}
		})
	}
}
