package days

import (
	"slices"

	"github.com/lordmoocow/aoc24/internal/utils"
)

type Day10 struct {
	topography utils.Grid[int]
	trails     map[utils.Point][][]utils.Point
}

func (d *Day10) Init(input string) {
	d.topography = utils.ParseIntGrid(input, utils.DefaultParseOptions)
}

func (d *Day10) PartOne() (result int) {
	d.getTrails(false)
	for _, t := range d.trails {
		result += len(t)
	}
	return result
}

func (d *Day10) getTrails(unique bool) {
	d.trails = map[utils.Point][][]utils.Point{}
	for xy := range d.topography.Find(0) {
		trail := make([]utils.Point, 1, 100)
		trail[0] = xy
		d.followTrail(trail, unique)
	}
}

func (d *Day10) followTrail(trail []utils.Point, unique bool) {
	current := trail[len(trail)-1]

	elevation, ok := d.topography.Get(current)
	if !ok {
		return
	}

	if elevation == 9 {
		d.addTrail(trail, unique)
		return
	}

	for next, nextElevation := range d.topography.NeighborsStraight(current) {
		if nextElevation-elevation == 1 {
			t := make([]utils.Point, len(trail)+1)
			copy(t, trail)
			t[len(t)-1] = next
			d.followTrail(t, unique)
		}
	}
}

func (d *Day10) addTrail(trail []utils.Point, unique bool) {
	if d.trails[trail[0]] == nil {
		d.trails[trail[0]] = make([][]utils.Point, 0, 100)
	}
	if !slices.ContainsFunc(d.trails[trail[0]], func(t []utils.Point) bool {
		if unique {
			return slices.Equal(t, trail)
		}
		return t[len(t)-1] == trail[len(trail)-1]
	}) {
		d.trails[trail[0]] = append(d.trails[trail[0]], trail)
	}
}

func (d *Day10) PartTwo() (result int) {
	d.getTrails(true)
	for _, t := range d.trails {
		result += len(t)
	}
	return result
}
