package days

import (
	"strconv"
	"strings"
	"unicode"
)

type Day3 struct {
	data string
}

func (d *Day3) Init(input string) {
	d.data = input
}

func (d *Day3) PartOne() int {
	return processInstructions(d.data)
}

func processInstructions(memory string) int {
	corruptInstructions := strings.Split(strings.TrimSpace(memory), "mul(")

	result := 0

outer:
	for _, ci := range corruptInstructions {
		fields := strings.Split(ci, ",")

		instruction := []int{}
		for _, f := range fields {
			// take first part if the instruction ends
			fs := strings.Split(f, ")")
			// check all characters of field are digits
			for _, c := range fs[0] {
				if !unicode.IsDigit(c) {
					continue outer
				}
			}
			i, _ := strconv.Atoi(fs[0])
			instruction = append(instruction, i)

			// if we did get a close then it's the end of this instruction
			if len(fs) > 1 {
				break
			}
		}
		result += mul(instruction)
	}

	return result
}

func mul(instruction []int) int {
	x := instruction[0]
	for _, y := range instruction[1:] {
		x *= y
	}
	return x
}

func (d *Day3) PartTwo() int {
	result := 0
	do := strings.Split(strings.TrimSpace(d.data), "do()")
	for _, d := range do {
		do := strings.SplitAfterN(d, "don't()", 2)[0]
		result += processInstructions(do)
	}
	return result
}
