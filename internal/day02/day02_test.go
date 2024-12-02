package day2

import (
	"testing"

	"github.com/lordmoocow/aoc24/assets"
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
