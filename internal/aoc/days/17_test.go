package days

import (
	"testing"
)

var day17 = []string{`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
	`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`}

func TestDaySeventeenPartOne(t *testing.T) {
	d := Day17{}
	d.Init(day17[0])

	want := "4,6,3,5,6,3,5,2,1,0"

	result := d.PartOne2()
	if result != want {
		t.Fatalf(`Day17.PartOne got %s but expected %s`, result, want)
	}
}

func BenchmarkDaySeventeenPartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day17{}
		d.Init(day17[0])
		d.PartOne()
	}
}
