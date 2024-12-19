package days

import (
	"testing"

	"github.com/lordmoocow/aoc24/internal/utils"
)

var day18 = []string{`5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`}

func TestDayEighteenPartOne(t *testing.T) {
	d := Day18{}
	d.Init(day18[0])
	d.memory = utils.NewGrid[rune](7, 7, '.')

	want := 22

	result := d.PartOne()
	if result != want {
		t.Fatalf(`Day18.PartOne got %d but expected %d`, result, want)
	}
}

func BenchmarkDayEighteenPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := Day18{}
		d.Init(day18[0])
		d.PartOne()
	}
}
