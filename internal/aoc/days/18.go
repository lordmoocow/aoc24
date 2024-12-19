package days

import (
	"fmt"
	"slices"
	"time"

	"github.com/gdamore/tcell"
	"github.com/lordmoocow/aoc24/internal/utils"
)

type Day18 struct {
	t      int
	bytes  []utils.Point
	memory utils.Grid[rune]
}

func (d *Day18) Init(input string) {
	d.t = 0
	d.bytes = utils.ParseCoordinates(input)
	d.memory = utils.NewGrid[rune](71, 71, ' ')
}

func (d *Day18) simulate(dt int) utils.Point {
	for range dt {
		if d.t >= len(d.bytes) {
			break
		}

		byte := d.bytes[d.t]
		d.memory.Set(byte, 'x')

		d.t++
	}
	return d.bytes[d.t-1]
}

type d18query struct {
	steps  int
	pos    utils.Point
	parent *d18query
}

func (d *Day18) path(start, finish utils.Point) []utils.Point {
	construct := func(q d18query) []utils.Point {
		path := make([]utils.Point, q.steps+1)
		s := q.steps + 1
		for {
			s--
			path[s] = q.pos
			if q.parent == nil {
				break
			}
			q = *q.parent
		}
		return path
	}
	h := func(p utils.Point) int {
		dx := p.X - finish.X
		dy := p.Y - finish.Y
		abs(&dx)
		abs(&dy)
		return dx + dy
	}
	visited := map[utils.Point]int{}

	queue := utils.NewPQ[d18query](100)
	queue.Push(0, d18query{0, start, nil})

	for {
		if queue.IsEmpty() {
			break
		}

		_, q := queue.Next()
		if q.pos == finish {
			return construct(q)
		}

		for next, content := range d.memory.NeighboursUDLR(q.pos) {
			if content != 'x' {
				steps := visited[q.pos] + 1
				if s, exists := visited[next]; !exists || steps < s {
					visited[next] = steps
					nq := d18query{steps, next, &q}
					priority := steps + h(next)
					queue.Push(priority, nq)
				}
			}
		}
	}

	return make([]utils.Point, 0)
}

func (d *Day18) PartOne() (result int) {
	screen, _ = tcell.NewScreen()
	if screen != nil {
		screen.Init()
		screen.Clear()
		defer func() {
			screen.Fini()
		}()
	}

	d.simulate(1024)
	path := d.path(utils.Point{X: 0, Y: 0}, d.memory.Bounds())

	if screen != nil {
	screenLoop:
		for {
			d.render()
			for _, p := range path {
				screen.SetContent(p.X, p.Y, '●', nil, tcell.StyleDefault)
			}
			screen.Show()
			//time.Sleep(100 * time.Millisecond)

			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					break screenLoop
				}
			}
		}
	}

	return len(path) - 1
}

func (d *Day18) PartTwo() (result int) {
	screen, _ = tcell.NewScreen()
	if screen != nil {
		screen.Init()
		screen.Clear()
	}

	d.simulate(1024)
	path := d.path(utils.Point{X: 0, Y: 0}, d.memory.Bounds())
	for {
		// if the byte corrupted our path then we need to reevaluate
		byte := d.simulate(1)
		if i := slices.Index(path, byte); i > -1 {
			// restart from step before this byte
			altPath := d.path(path[i-1], d.memory.Bounds())
			// if we got no route then this is the end :(
			if len(altPath) == 0 {
				defer fmt.Printf("%d,%d", byte.X, byte.Y)
				break
			}
			// rebuild path with alternative route
			path = append(path[:i-1], altPath...)
		}

		if screen != nil {
			d.render()
			for _, p := range path {
				screen.SetContent(p.X, p.Y, '●', nil, tcell.StyleDefault)
			}
			screen.Show()
			time.Sleep(1 * time.Millisecond)
		}
	}
	if screen != nil {
		screen.Fini()
	}
	return result
}

func (d *Day18) render() {
	for p, cell := range d.memory.All() {
		screen.SetContent(p.X, p.Y, cell, nil, tcell.StyleDefault)
	}
}
