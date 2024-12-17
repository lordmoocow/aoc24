package days

import (
	"github.com/lordmoocow/aoc24/internal/utils"
)

type Day16 struct {
	maze       utils.Grid[rune]
	reindeer   reindeer
	start, end utils.Point
}

type reindeer struct {
	pos utils.Point
	dir utils.Vec
}

func (d *Day16) Init(input string) {
	d.maze = utils.ParseRuneGrid(input, utils.DefaultParseOptions)
	for s := range d.maze.Find('S') {
		d.start = s
		d.reindeer = reindeer{s, utils.Right}
	}
	for e := range d.maze.Find('E') {
		d.end = e
	}
}

type mazequery struct {
	reindeer
	steps, turns int
}

func (d *Day16) search() (result int) {
	visited := map[utils.Point]bool{}
	queue := utils.NewPQ[mazequery](100)
	queue.Push(0, mazequery{d.reindeer, 0, 0})

	for {
		if queue.IsEmpty() {
			break
		}

		cost, q := queue.Next()
		if q.pos == d.end {
			return cost
		}

		for _, dir := range []utils.Vec{
			utils.Up,
			utils.Down,
			utils.Left,
			utils.Right,
		} {
			p := q.pos.Add(&dir)
			if !visited[p] {
				visited[p] = true
				if v, ok := d.maze.Get(p); ok && v != '#' {
					r := q
					r.pos = p
					r.dir = dir
					r.steps++
					c := cost + 1
					if q.dir != dir {
						c += 1000
						r.turns++
					}
					queue.Push(c, r)
				}
			}
		}
	}

	return 0
}

func (d *Day16) PartOne() (result int) {
	return d.search()
}

func (d *Day16) PartTwo() (result int) {
	return result
}
