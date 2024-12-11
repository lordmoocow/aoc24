package days

import (
	"testing"
)

var day7 = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestDaySevenPartOne(t *testing.T) {
	d := Day7{}
	d.Init(day7)

	want := 3749

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day7.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDaySevenPartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day7{}
		d.Init(day7)
		d.PartOne()
	}
}

func TestDaySevenPartTwo(t *testing.T) {
	d := Day7{}
	d.Init(day7)

	want := 11387
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day7.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDaySevenPartTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day7{}
		d.Init(day7)
		d.PartTwo()
	}
}
