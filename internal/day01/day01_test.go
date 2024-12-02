package day1

import (
	"testing"

	"github.com/lordmoocow/aoc24/assets"
)

func TestDayOnePartOne(t *testing.T) {
	d := Day1{}
	d.Init(assets.TestData(1))

	want := 11
	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day1.PartOne = %d, want match for %d`, result, want)
	}
}

func TestDayOnePartTwo(t *testing.T) {
	d := Day1{}
	d.Init(assets.TestData(1))

	want := 31
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day1.PartOne = %d, want match for %d`, result, want)
	}
}
