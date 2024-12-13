package days

import (
	"slices"

	"github.com/lordmoocow/aoc24/internal/utils"
)

type Day12 struct {
	garden utils.Grid[rune]
}

type region struct {
	plots  []utils.Point
	fences int
}

func (r *region) price() int {
	return len(r.plots) * r.fences
}

func (d *Day12) Init(input string) {
	d.garden = utils.ParseRuneGrid(input, utils.DefaultParseOptions)
}

func (d *Day12) crops() map[rune][]region {
	crops := map[rune][]region{}
	visited := map[utils.Point]bool{}
	for p, crop := range d.garden.All() {
		if visited[p] {
			continue
		}
		c, exists := crops[crop]
		if !exists {
			c = make([]region, 0, 10)
		}

		r := region{}
		r.plots = append(r.plots, p)
		visited[p] = true

		queue := make([]utils.Point, 1, 10)
		queue[0] = p
		for {
			if len(queue) == 0 {
				break
			}
			pop := queue[len(queue)-1]
			queue = queue[:len(queue)-1]

			edges := 4 // account for lack of neighbours at edges
			for n, ncrop := range d.garden.NeighboursUDLR(pop) {
				edges--
				if crop != ncrop {
					r.fences++
					continue
				}

				if !slices.Contains(r.plots, n) {
					queue = append(queue, n)
					r.plots = append(r.plots, n)
					visited[n] = true
				}
			}
			r.fences += edges
		}

		c = append(c, r)
		crops[crop] = c
	}
	return crops
}

func (d *Day12) PartOne() (result int) {
	crops := d.crops()
	for _, crop := range crops {
		for _, region := range crop {
			result += region.price()
		}
	}
	return result
}

func (d *Day12) PartTwo() (result int) {
	return result
}
