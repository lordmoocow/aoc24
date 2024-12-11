package days

import (
	"testing"
)

var day2 = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestDayTwoPartOne(t *testing.T) {
	d := Day2{}
	d.Init(day2)

	want := 2
	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day2.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayTwoPartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day2{}
		d.Init(day2)
		d.PartOne()
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	d := Day2{}
	d.Init(day2)

	want := 4
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day2.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayTwoPartTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day2{}
		d.Init(day2)
		d.PartTwo()
	}
}
