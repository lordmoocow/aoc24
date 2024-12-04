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
