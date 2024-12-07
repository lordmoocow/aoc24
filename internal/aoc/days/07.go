package days

import (
	"strconv"
	"strings"
)

type Day7 struct {
	calibrations []calibration
}

type calibration struct {
	result int
	values []int
}

func (d *Day7) Init(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	d.calibrations = make([]calibration, len(lines))
	for i, line := range lines {
		f := strings.Fields(line)
		c := calibration{}
		c.result, _ = strconv.Atoi(f[0][:len(f[0])-1])
		vals := f[1:]
		c.values = make([]int, len(vals))
		for i, val := range vals {
			c.values[i], _ = strconv.Atoi(val)
		}
		d.calibrations[i] = c
	}
}

func (d *Day7) PartOne() (result int) {
	for _, c := range d.calibrations {
		if ok := c.solve(c.values[0], c.values[1:]); ok {
			result += c.result
		}
	}
	return result
}

type operator = func(a, b int) int

var ops [2]operator = [2]operator{
	add,
	mul,
}

func add(v int, n int) int {
	return v + n
}

func mul(v int, n int) int {
	return v * n
}

func (c calibration) solve(x int, equation []int) bool {
	if x < c.result && len(equation) > 0 {
		for _, op := range ops {
			o := op(x, equation[0])
			if c.solve(o, equation[1:]) {
				return true
			}
		}
	}
	return x == c.result
}

func (d *Day7) PartTwo() (result int) {
	return result
}