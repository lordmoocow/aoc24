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
	if x < 0 || y < 0 || x >= len(d.wordSearch[0]) || y >= len(d.wordSearch) {
		return false
	}
	v := d.wordSearch[y][x]
	if rune(xmas[i]) == v {
		nx, ny := x+xdir, y+ydir
		return d.search(i+1, nx, ny, xdir, ydir)
	}
	return false
}

func (d *Day4) PartTwo() int {
	return 0
}
