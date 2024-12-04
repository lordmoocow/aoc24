package days

import (
	"testing"

	"github.com/lordmoocow/aoc24/internal/assets"
)

func TestDayFourPartOne(t *testing.T) {
	d := Day4{}
	d.Init(assets.TestData(4))

	want := 18

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day4.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayFourPartOne(b *testing.B) {
	d := Day4{}
	d.Init(assets.InputData(4))

	for i := 0; i < b.N; i++ {
		d.PartOne()
	}
}

func TestDayFourPartTwo(t *testing.T) {
	d := Day4{}
	d.Init(assets.TestData(4))

	want := 9
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day4.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayFourPartTwo(b *testing.B) {
	d := Day4{}
	d.Init(assets.InputData(4))

	for i := 0; i < b.N; i++ {
		d.PartTwo()
	}
}
