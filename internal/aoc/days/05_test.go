package days

import (
	"testing"
)

var day5 = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

func TestDayFivePartOne(t *testing.T) {
	d := Day5{}
	d.Init(day5)

	want := 143

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day5.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayFivePartOne(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day5{}
		d.Init(day5)
		d.PartOne()
	}
}

func TestDayFivePartTwo(t *testing.T) {
	d := Day5{}
	d.Init(day5)

	want := 123
	result := d.PartTwo()
	if result != want {
		t.Fatalf(`Day5.PartTwo got %d but expected %d`, result, want)
	}
}

func BenchmarkDayFivePartTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := Day5{}
		d.Init(day5)
		d.PartTwo()
	}
}
