package days

import (
	"testing"
)

var day16 = []string{`###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`,
	`#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`}

func TestDaySixteenPartOne(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{expected: 7036, input: day16[0]},
		{expected: 11048, input: day16[1]},
	}
	for _, test := range tests {
		d := Day16{}
		d.Init(test.input)

		result := d.PartOne()
		if result != test.expected {
			t.Errorf(`Day16.PartOne got %d but expected %d`, result, test.expected)
		}
	}
}

func BenchmarkDaySixteenPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := Day16{}
		d.Init(day16[0])
		d.PartOne()
	}
}
