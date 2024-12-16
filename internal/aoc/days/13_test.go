package days

import (
	"testing"
)

var day13 = []string{`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`,

	`Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176`,

	`Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450`,

	`Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
}

func TestDayThirteenPartOne(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{expected: 280, input: day13[0]},
		{expected: 0, input: day13[1]},
		{expected: 200, input: day13[2]},
		{expected: 0, input: day13[3]},
	}
	for _, test := range tests {
		d := Day13{}
		d.Init(test.input)

		result := d.PartOne()
		if result != test.expected {
			t.Errorf(`Day13.PartOne got %d but expected %d`, result, test.expected)
		}
	}
}

func BenchmarkDayThirteenPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := Day13{}
		d.Init(day13[0])
		d.PartOne()
	}
}
