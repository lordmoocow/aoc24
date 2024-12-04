package days

import (
	"strconv"
	"strings"
)

type Day2 struct {
	reports [][]int
	scratch [10]int
}

func (d *Day2) Init(input string) {
	// Split input by line
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Set known capacity
	d.reports = make([][]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		d.reports[i] = make([]int, len(fields))
		if len(fields) > 0 {
			for j, reactorLevel := range fields {
				d.reports[i][j], _ = strconv.Atoi(reactorLevel)
			}
		}
	}
}

func (d *Day2) PartOne() (safeReports int) {
	for _, report := range d.reports {
		if d.checkSafe(report) {
			safeReports++
		}
	}
	return safeReports
}

func (d *Day2) checkSafe(report []int) bool {
	increasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// are we going in the wrong direction?
		if (increasing && diff < 0) || (!increasing && diff > 0) {
			return false
		}

		// does it exceed threshold?
		abs(&diff)
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func abs(x *int) {
	if *x < 0 {
		*x *= -1
	}
}

func (d *Day2) PartTwo() (safeReports int) {
	for _, report := range d.reports {
		if d.checkSafeWithDampener(report) {
			safeReports++
		}
	}
	return safeReports
}

func (d *Day2) checkSafeWithDampener(report []int) bool {
	for i := range report {
		// copy slice before and after i into scratch space
		// this "removes" the value at i from the report
		copy(d.scratch[:i], report[:i])
		copy(d.scratch[i:], report[i+1:])
		damped := d.scratch[:len(report)-1]

		// check the updated report is safe
		if d.checkSafe(damped) {
			return true
		}
	}
	return false
}
