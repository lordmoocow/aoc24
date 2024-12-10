package days

import (
	"testing"
)

var day1 = `3   4
4   3
2   5
1   3
3   9
3   3
`

func TestDayOnePartOne(t *testing.T) {
	d := Day1{}
	d.Init(day1)

	want := 11
	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day1.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayOnePartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day1{}
		d.Init(day1)
		d.PartOne()
	}
}

func TestDayOnePartTwo(t *testing.T) {
	d := Day1{}
	d.Init(day1)

	want := 31
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day1.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayOnePartTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day1{}
		d.Init(day1)
		d.PartTwo()
	}
}
