package day1

import (
	"testing"

	"github.com/lordmoocow/aoc24/assets"
)

func TestDayOnePartOne(t *testing.T) {
	d := Day1{}
	d.Init(assets.TestData(1))

	want := 11
	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day1.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayOnePartOne(b *testing.B) {
	d := Day1{}
	d.Init(assets.InputData(1))

	for i := 0; i < b.N; i++ {
		d.PartOne()
	}
}

func TestDayOnePartTwo(t *testing.T) {
	d := Day1{}
	d.Init(assets.TestData(1))

	want := 31
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day1.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayOnePartTwo(b *testing.B) {
	d := Day1{}
	d.Init(assets.InputData(1))

	for i := 0; i < b.N; i++ {
		d.PartTwo()
	}
}
