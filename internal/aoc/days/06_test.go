package days

import (
	"testing"
)

var day6 = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

func TestDaySixPartOne(t *testing.T) {
	d := Day6{}
	d.Init(day6)

	want := 41

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day6.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDaySixPartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day6{}
		d.Init(day6)
		d.PartOne()
	}
}

func TestDaySixPartTwo(t *testing.T) {
	d := Day6{}
	d.Init(day6)

	want := 6
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day6.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDaySixPartTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day6{}
		d.Init(day6)
		d.PartTwo()
	}
}
