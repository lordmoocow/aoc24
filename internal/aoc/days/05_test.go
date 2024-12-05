package days

import (
	"testing"

	"github.com/lordmoocow/aoc24/internal/assets"
)

func TestDayFivePartOne(t *testing.T) {
	d := Day5{}
	d.Init(assets.TestData(5))

	want := 143

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day5.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayFivePartOne(b *testing.B) {
	d := Day5{}
	d.Init(assets.InputData(5))

	for i := 0; i < b.N; i++ {
		d.PartOne()
	}
}

func TestDayFivePartTwo(t *testing.T) {
	d := Day5{}
	d.Init(assets.TestData(5))

	want := 123
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day5.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayFivePartTwo(b *testing.B) {
	d := Day5{}
	d.Init(assets.InputData(5))

	for i := 0; i < b.N; i++ {
		d.PartTwo()
	}
}
