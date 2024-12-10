package days

import (
	"testing"
)

var day8 = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

func TestDayEightPartOne(t *testing.T) {
	d := Day8{}
	d.Init(day8)

	want := 14

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day8.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayEightPartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day8{}
		d.Init(day8)
		d.PartOne()
	}
}

func TestDayEightPartTwo(t *testing.T) {
	d := Day8{}
	d.Init(day8)

	want := 34

	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day8.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayEightPartTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day8{}
		d.Init(day8)
		d.PartTwo()
	}
}
