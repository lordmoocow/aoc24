package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day17 struct {
	A, B, C int
	program []int
}

type computer struct {
	A, B, C int
	i       int
}

func (c *computer) exec(program []int) []int {
	output := make([]int, 0, 10)
	for {
		// halt if end of program
		if c.i >= len(program) {
			break
		}

		operand := 0
		if c.i+1 < len(program) {
			operand = program[c.i+1]
		}

		switch program[c.i] {
		case 0:
			c.adv(operand)
		case 1:
			c.bxl(operand)
		case 2:
			c.bst(operand)
		case 3:
			if c.jnz(operand) {
				continue
			}
		case 4:
			c.bxc()
		case 5:
			if o, ok := c.out(operand); ok {
				output = append(output, o)
			}
		case 6:
			c.bdv(operand)
		case 7:
			c.cdv(operand)
		}
		c.i += 2
	}
	return output
}

func (c *computer) combo(operand int) (int, bool) {
	if operand <= 3 {
		return operand, true
	}

	switch operand {
	case 4:
		return c.A, true
	case 5:
		return c.B, true
	case 6:
		return c.C, true
	}

	return 0, false
}
func (c *computer) adv(combo int) { c.dv(&c.A, combo) }
func (c *computer) bdv(combo int) { c.dv(&c.B, combo) }
func (c *computer) cdv(combo int) { c.dv(&c.C, combo) }
func (c *computer) dv(register *int, combo int) {
	if o, ok := c.combo(combo); ok {
		*register = int(float64(c.A) / math.Pow(2.0, float64(o)))
	}
}
func (c *computer) bxl(literal int) {
	c.B ^= literal
}
func (c *computer) bst(combo int) {
	if o, ok := c.combo(combo); ok {
		c.B = o % 8
	}
}
func (c *computer) jnz(literal int) bool {
	if c.A == 0 {
		return false
	}
	c.i = literal
	return true
}
func (c *computer) bxc() {
	c.B ^= c.C
}
func (c *computer) out(combo int) (int, bool) {
	if o, ok := c.combo(combo); ok {
		return o % 8, true
	}
	return 0, false
}

func (d *Day17) Init(input string) {
	s, p, _ := strings.Cut(strings.TrimSpace(input), "\n\n")
	state := strings.Split(s, "\n")
	d.A, _ = strconv.Atoi(state[0][12:])
	d.B, _ = strconv.Atoi(state[1][12:])
	d.C, _ = strconv.Atoi(state[2][12:])
	program := strings.Split(p[9:], ",")
	d.program = make([]int, len(program))
	for i, p := range program {
		d.program[i], _ = strconv.Atoi(p)
	}
}

func (d *Day17) PartOne2() string {
	c := computer{A: d.A, B: d.B, C: d.C}
	out := c.exec(d.program)
	csv := strings.Replace(fmt.Sprintf("%#v", out), " ", "", -1)
	return csv[6 : len(csv)-1]
}
func (d *Day17) PartOne() (result int) {
	s := d.PartOne2()
	fmt.Printf("%s\n", s)
	return result
}

func (d *Day17) PartTwo() (result int) {
	return result
}
