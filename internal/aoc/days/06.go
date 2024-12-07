package days

import "strings"

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
	guard guard
}

func (d *Day6) Init(input string) {
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
	result = 1
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
		if v == '.' {
			distance++
		}
		d.guard.pos.x += d.guard.dir.x
		d.guard.pos.y += d.guard.dir.y
		d.lab[d.guard.pos.y][d.guard.pos.x] = 'x'

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

	return false, distance
}

func (d *Day6) PartTwo() (result int) {
	return result
}
