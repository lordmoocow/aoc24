package days

import (
	"strconv"
	"strings"
	"unicode"
)

type Day3 struct {
	data    string
	scratch [2]int
}

func (d *Day3) Init(input string) {
	d.data = input
}

func (d *Day3) PartOne() int {
	return d.processInstructions(d.data)
}

func (d *Day3) processInstructions(memory string) (result int) {
	// if we split on mul( we know that the start of every element is a potential instruction
	instructions := strings.Split(strings.TrimSpace(memory), "mul(")

scanInstructions:
	for _, ci := range instructions {
		// we know we have a valid start, so just find the end and discard the rest
		instructionBody, _, valid := strings.Cut(ci, ")")
		// if it never ends then it can't be valid
		if !valid {
			continue
		}
		// split out param values to examine
		rawParams := strings.Split(instructionBody, ",")

		for i, p := range rawParams {
			// check all characters are digits
			for _, c := range p {
				// if we find anything invalid at this point we know this instruction is corrupt
				// move on to the next one
				if !unicode.IsDigit(c) {
					continue scanInstructions
				}
			}
			// save the instruction parameter
			d.scratch[i], _ = strconv.Atoi(p)
		}
		// execute valid mul instruction
		result += d.mul(d.scratch[:len(rawParams)])
	}

	return result
}

func (d *Day3) mul(instruction []int) int {
	x := instruction[0]
	for _, y := range instruction[1:] {
		x *= y
	}
	return x
}

func (d *Day3) PartTwo() int {
	result := 0
	do := strings.Split(strings.TrimSpace(d.data), "do()")
	for _, x := range do {
		do := strings.SplitAfterN(x, "don't()", 2)[0]
		result += d.processInstructions(do)
	}
	return result
}
