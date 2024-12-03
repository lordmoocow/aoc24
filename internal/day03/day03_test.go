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
