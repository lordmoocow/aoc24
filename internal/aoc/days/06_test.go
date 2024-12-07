package days

import (
	"testing"

	"github.com/lordmoocow/aoc24/internal/assets"
)

func TestDaySixPartOne(t *testing.T) {
	d := Day6{}
	d.Init(assets.TestData(6))

	want := 41

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day6.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDaySixPartOne(b *testing.B) {
	d := Day6{}
	d.Init(assets.InputData(6))

	for i := 0; i < b.N; i++ {
		d.PartOne()
	}
}
