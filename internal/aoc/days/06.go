package days

import (
	"slices"
	"strings"
)

type vec struct {
	x, y int
}

var (
	up    = vec{0, -1}
	down  = vec{0, 1}
	left  = vec{-1, 0}
	right = vec{1, 0}
)

type guard struct {
	pos, dir vec
}

func (g *guard) turn() {
	switch g.dir {
	case up:
		g.dir = right
	case right:
		g.dir = down
	case down:
		g.dir = left
	case left:
		g.dir = up
	}
}

type Day6 struct {
	lab   [][]rune
	path  map[vec]map[vec]bool
	guard guard
}

func (d *Day6) Init(input string) {
	d.path = map[vec]map[vec]bool{}
	rows := strings.Split(strings.TrimSpace(input), "\n")
	d.lab = make([][]rune, len(rows))
	for y, row := range rows {
		if x := strings.IndexRune(row, '^'); x > -1 {
			d.guard = guard{pos: vec{x, y}, dir: up}
		}
		d.lab[y] = []rune(row)
	}
}

func (d *Day6) PartOne() (result int) {
	for {
		blocked, dist := d.move()
		result += dist
		if !blocked {
			return result
		}
		d.guard.turn()
	}
}

func (d *Day6) move() (blocked bool, distance int) {
	for {
		v := d.lab[d.guard.pos.y+d.guard.dir.y][d.guard.pos.x+d.guard.dir.x]
		if v == '#' {
			return true, distance
		}
		if d.path[d.guard.pos] == nil {
			distance++
			d.path[d.guard.pos] = map[vec]bool{}
		}
		d.path[d.guard.pos][d.guard.dir] = true
		d.guard.pos.x += d.guard.dir.x
		d.guard.pos.y += d.guard.dir.y

		if d.guard.dir.x > 0 && d.guard.pos.x == len(d.lab[d.guard.pos.y])-1 {
			break
		}
		if d.guard.dir.y > 0 && d.guard.pos.y == len(d.lab)-1 {
			break
		}
		if d.guard.dir.x < 0 && d.guard.pos.x == 0 {
			break
		}
		if d.guard.dir.y < 0 && d.guard.pos.y == 0 {
			break
		}
	}

	// catch the last space
	if d.path[d.guard.pos] == nil {
		distance++
		d.path[d.guard.pos] = map[vec]bool{}
	}
	d.path[d.guard.pos][d.guard.dir] = true
	return false, distance
}

func (d *Day6) PartTwo() (result int) {
	clone := d.clone()
	// calculate initial guard path
	clone.PartOne()

	for pos := range clone.path {
		// skip initial position
		if pos == d.guard.pos {
			continue
		}
		// fresh clone to run again
		c := d.clone()

		// add new obstruction in path
		c.lab[pos.y][pos.x] = '#'

		// run guard movement
		for {
			if blocked, _ := c.move(); !blocked {
				break
			}

			c.guard.turn()

			// have we been this way already?
			if c.path[c.guard.pos][c.guard.dir] {
				d.lab[pos.y][pos.x] = 'O'
				result++
				break
			}
		}
		// remove the obstacle
		c.lab[pos.y][pos.x] = '.'
	}
	return result
}

func (d Day6) clone() Day6 {
	d.path = map[vec]map[vec]bool{}
	d.lab = slices.Clone(d.lab)
	for i, row := range d.lab {
		d.lab[i] = slices.Clone(row)
	}
	return d
}
