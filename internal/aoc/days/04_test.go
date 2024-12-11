package days

import (
	"testing"
)

var day4 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func TestDayFourPartOne(t *testing.T) {
	d := Day4{}
	d.Init(day4)

	want := 18

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day4.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayFourPartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day4{}
		d.Init(day4)
		d.PartOne()
	}
}

func TestDayFourPartTwo(t *testing.T) {
	d := Day4{}
	d.Init(day4)

	want := 9
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day4.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayFourPartTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day4{}
		d.Init(day4)
		d.PartTwo()
	}
}
