package days

import (
	"testing"
)

var day9 = `2333133121414131402`

func TestDayNinePartOne(t *testing.T) {
	d := Day9{}
	d.Init(day9)

	want := 1928

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day9.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayNinePartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day9{}
		d.Init(day9)
		d.PartOne()
	}
}

func TestDayNinePartTwo(t *testing.T) {
	d := Day9{}
	d.Init(day9)

	want := 2858

	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day9.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayNinePartTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day9{}
		d.Init(day9)
		d.PartTwo()
	}
}
