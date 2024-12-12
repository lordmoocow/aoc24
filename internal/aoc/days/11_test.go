package days

import (
	"testing"
)

var day11 = `125 17`

func TestDayElevenPartOne(t *testing.T) {
	d := Day11{}
	d.Init(day11)

	want := 55312

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day11.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayElevenPartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day11{}
		d.Init(day11)
		d.PartOne()
	}
}
