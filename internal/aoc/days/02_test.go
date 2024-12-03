package days

import (
	"testing"

	"github.com/lordmoocow/aoc24/internal/assets"
)

func TestDayTwoPartOne(t *testing.T) {
	d := Day2{}
	d.Init(assets.TestData(2))

	want := 2
	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day2.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayTwoPartOne(b *testing.B) {
	d := Day2{}
	d.Init(assets.InputData(2))

	for i := 0; i < b.N; i++ {
		d.PartOne()
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	d := Day2{}
	d.Init(assets.TestData(2))

	want := 4
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day2.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayTwoPartTwo(b *testing.B) {
	d := Day2{}
	d.Init(assets.InputData(2))

	for i := 0; i < b.N; i++ {
		d.PartTwo()
	}
}
