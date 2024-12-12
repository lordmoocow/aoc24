package aoc

import (
	"errors"

	"github.com/lordmoocow/aoc24/internal/aoc/days"
)

type Day interface {
	Init(string)
	PartOne() (result int)
	PartTwo() (result int)
}

func Run(day, part int, input string) (result int, err error) {
	if day < 1 || day > 25 {
		err = errors.New("day must be within inclusive range 1-25")
	} else if part != 1 && part != 2 {
		err = errors.New("part must be 1 or 2")
	}
	if err != nil {
		return 0, err
	}

	d, err := getDay(day)
	if err != nil {
		return 0, err
	}

	d.Init(input)

	switch part {
	case 1:
		result = d.PartOne()
	case 2:
		result = d.PartTwo()
	}

	return result, err
}

func getDay(day int) (d Day, err error) {
	switch day {
	case 1:
		d = &days.Day1{}
	case 2:
		d = &days.Day2{}
	case 3:
		d = &days.Day3{}
	case 4:
		d = &days.Day4{}
	case 5:
		d = &days.Day5{}
	case 6:
		d = &days.Day6{}
	case 7:
		d = &days.Day7{}
	case 8:
		d = &days.Day8{}
	case 9:
		d = &days.Day9{}
	case 10:
		d = &days.Day10{}
	case 11:
		d = &days.Day11{}

	default:
		err = errors.New("not implemented")
	}

	return d, err
}
