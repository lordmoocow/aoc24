package days

import (
	"testing"

	"github.com/lordmoocow/aoc24/internal/assets"
)

func TestDaySevenPartOne(t *testing.T) {
	d := Day7{}
	d.Init(assets.TestData(7))

	want := 3749

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day7.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDaySevenPartOne(b *testing.B) {
	d := Day7{}
	d.Init(assets.InputData(7))

	for i := 0; i < b.N; i++ {
		d.PartOne()
	}
}

func TestDaySevenPartTwo(t *testing.T) {
	d := Day7{}
	d.Init(assets.TestData(7))

	want := 11387
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day7.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDaySevenPartTwo(b *testing.B) {
	d := Day7{}
	d.Init(assets.InputData(7))

	for i := 0; i < b.N; i++ {
		d.PartTwo()
	}
}
