package day1

import (
	_ "embed"
	"testing"
)

//go:embed test
var testinput string

func TestPartOne(t *testing.T) {
	input = testinput

	d := Day1{}
	d.Init()

	want := 11
	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day1.PartOne = %d, want match for %d`, result, want)
	}
}

func TestPartTwo(t *testing.T) {
	input = testinput

	d := Day1{}
	d.Init()

	want := 31
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day1.PartOne = %d, want match for %d`, result, want)
	}
}
