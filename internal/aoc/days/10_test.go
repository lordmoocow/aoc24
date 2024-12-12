package days

import (
	"testing"
)

var day10 = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestDayTenPartOne(t *testing.T) {
	d := Day10{}
	d.Init(day10)

	want := 36

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day10.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayTenPartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day10{}
		d.Init(day10)
		d.PartOne()
	}
}

func TestDayTenPartTwo(t *testing.T) {
	d := Day10{}
	d.Init(day10)

	want := 81

	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day10.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayTenPartTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day10{}
		d.Init(day10)
		d.PartTwo()
	}
}
