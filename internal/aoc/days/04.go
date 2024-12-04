package days

import (
	"strings"
)

type Day4 struct {
	wordSearch [][]rune
}

func (d *Day4) Init(input string) {
	rows := strings.Split(input, "\n")
	d.wordSearch = make([][]rune, len(rows)-1)
	for i, row := range rows {
		if len(row) > 0 {
			d.wordSearch[i] = []rune(row)
		}
	}
}

func (d *Day4) PartOne() (result int) {
	for y, row := range d.wordSearch {
		for x, v := range row {
			if v == rune("X"[0]) {
				if d.search(0, x, y, -1, -1) {
					result++
				}
				if d.search(0, x, y, -1, 0) {
					result++
				}
				if d.search(0, x, y, 0, -1) {
					result++
				}
				if d.search(0, x, y, 1, 0) {
					result++
				}
				if d.search(0, x, y, 0, 1) {
					result++
				}
				if d.search(0, x, y, 1, -1) {
					result++
				}
				if d.search(0, x, y, -1, 1) {
					result++
				}
				if d.search(0, x, y, 1, 1) {
					result++
				}
			}
		}
	}
	return result
}

func (d *Day4) search(i, x, y, xdir, ydir int) bool {
	const xmas string = "XMAS"
	if i == len(xmas) {
		return true
	}

	if d.check(rune(xmas[i]), x, y) {
		nx, ny := x+xdir, y+ydir
		return d.search(i+1, nx, ny, xdir, ydir)
	}
	return false
}

func (d *Day4) check(want rune, x, y int) bool {
	if x < 0 || y < 0 || x >= len(d.wordSearch[0]) || y >= len(d.wordSearch) {
		return false
	}
	v := d.wordSearch[y][x]
	return want == v
}

func (d *Day4) PartTwo() (result int) {
	for y, row := range d.wordSearch {
		for x, v := range row {
			if v == 'A' {
				if (d.check('M', x-1, y-1) && d.check('S', x+1, y+1)) || (d.check('S', x-1, y-1) && d.check('M', x+1, y+1)) {
					if (d.check('M', x-1, y+1) && d.check('S', x+1, y-1)) || (d.check('S', x-1, y+1) && d.check('M', x+1, y-1)) {
						result++
					}
				}
			}
		}
	}
	return result
}
