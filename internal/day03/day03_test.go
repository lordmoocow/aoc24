package day3

import (
	"testing"

	"github.com/lordmoocow/aoc24/assets"
)

func TestDayThreePartOne(t *testing.T) {
	d := Day3{}
	d.Init(assets.TestData(3))

	want := 161

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day3.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayThreePartOne(b *testing.B) {
	d := Day3{}
	d.Init(assets.InputData(3))

	for i := 0; i < b.N; i++ {
		d.PartOne()
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	d := Day3{}
	d.Init("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

	want := 48
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day3.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayThreePartTwo(b *testing.B) {
	d := Day3{}
	d.Init(assets.InputData(3))

	for i := 0; i < b.N; i++ {
		d.PartTwo()
	}
}
