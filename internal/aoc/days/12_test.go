package days

import (
	"testing"
)

var day12 = []string{`AAAA
BBCD
BBCC
EEEC`,
	`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
	`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
}

func TestDayTwelvePartOne(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{expected: 140, input: day12[0]},
		{expected: 772, input: day12[1]},
		{expected: 1930, input: day12[2]},
	}
	for _, test := range tests {
		d := Day12{}
		d.Init(test.input)

		result := d.PartOne()
		if result != test.expected {
			t.Errorf(`Day12.PartOne got %d but expected %d`, result, test.expected)
		}
	}
}

func BenchmarkDayTwelvePartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := Day12{}
		d.Init(day12[0])
		d.PartOne()
	}
}
