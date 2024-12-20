package days

import (
	"slices"
	"strconv"
	"strings"
)

type Day1 struct {
	left  []int
	right []int
}

func (d *Day1) Init(input string) {
	// Split input by line
	lines := strings.Split(input, "\n")

	// Set known capacity
	d.left = make([]int, len(lines))
	d.right = make([]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 0 {
			d.left[i], _ = strconv.Atoi(fields[0])
			d.right[i], _ = strconv.Atoi(fields[1])
		}
	}
}

func (d *Day1) PartOne() (result int) {
	slices.Sort(d.left)
	slices.Sort(d.right)

	for i := 0; i < len(d.left); i++ {
		distance := d.left[i] - d.right[i]
		if distance < 0 {
			distance *= -1
		}
		result += distance
	}

	return result
}

func (d *Day1) PartTwo() (result int) {
	slices.Sort(d.left)
	slices.Sort(d.right)

	right := d.right[:]
	count := 0
	for i, x := range d.left {

		if count == 0 {
			for _, y := range right {
				if y > x {
					break
				} else if y < x {
					continue
				}
				count++
			}
		}

		result += x * count

		if i < len(d.left)-1 && x != d.left[i+1] {
			if count < len(right) {
				right = right[count:]
			}
			count = 0
		}
	}

	return result
}
