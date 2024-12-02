package aoc

type Day interface {
	Init(string)
	PartOne() int
	PartTwo() int
}
