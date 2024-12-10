package days

import (
	"testing"

	"github.com/lordmoocow/aoc24/internal/assets"
)

func TestDayNinePartOne(t *testing.T) {
	d := Day9{}
	d.Init(assets.TestData(9))

	want := 1928

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day9.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayNinePartOne(b *testing.B) {
	d := Day9{}
	d.Init(assets.InputData(9))

	for i := 0; i < b.N; i++ {
		d.PartOne()
	}
}

func TestDayNinePartTwo(t *testing.T) {
	d := Day9{}
	d.Init(assets.TestData(9))

	want := 2858

	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day9.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayNinePartTwo(b *testing.B) {
	d := Day9{}
	d.Init(assets.InputData(9))

	for i := 0; i < b.N; i++ {
		d.PartTwo()
	}
}
