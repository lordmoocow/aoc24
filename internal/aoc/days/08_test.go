package days

import (
	"testing"

	"github.com/lordmoocow/aoc24/internal/assets"
)

func TestDayEightPartOne(t *testing.T) {
	d := Day8{}
	d.Init(assets.TestData(8))

	want := 14

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day8.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayEightPartOne(b *testing.B) {
	d := Day8{}
	d.Init(assets.InputData(8))

	for i := 0; i < b.N; i++ {
		d.PartOne()
	}
}
