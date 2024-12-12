package days

import (
	"slices"
	"strconv"
	"strings"
)

type Day10 struct {
	topography    [][]int
	width, height int
	trails        map[vec][][]vec
}

func (d *Day10) Init(input string) {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	d.height = len(rows) - 1
	d.topography = make([][]int, d.height+1)
	for i, row := range rows {
		if len(row) > 0 {
			d.width = len(row) - 1
			d.topography[i] = make([]int, d.width+1)
			for j, v := range row {
				x, _ := strconv.Atoi(string(v))
				d.topography[i][j] = x
			}
		}
	}
}

func (d *Day10) PartOne() (result int) {
	d.getTrails(false)
	for _, t := range d.trails {
		result += len(t)
	}
	return result
}

func (d *Day10) getTrails(unique bool) {
	d.trails = map[vec][][]vec{}
	for y, row := range d.topography {
		for x, elevation := range row {
			if elevation == 0 {
				xy := vec{x, y}
				trail := make([]vec, 1, 100)
				trail[0] = xy
				d.followTrail(trail, unique)
			}
		}
	}
}

func (d *Day10) followTrail(trail []vec, unique bool) {
	current := trail[len(trail)-1]

	elevation := d.topography[current.y][current.x]

	if elevation == 9 {
		d.addTrail(trail, unique)
		return
	}

	for _, dir := range []vec{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	} {
		next := vec{current.x + dir.x, current.y + dir.y}
		if d.inbounds(next) {
			nextElevation := d.topography[next.y][next.x]
			if nextElevation-elevation == 1 {
				t := make([]vec, len(trail)+1)
				copy(t, trail)
				t[len(t)-1] = next
				d.followTrail(t, unique)
			}
		}
	}
}

func (d *Day10) inbounds(point vec) bool {
	if point.x < 0 || point.y < 0 || point.x > d.width || point.y > d.height {
		return false
	}
	return true
}

func (d *Day10) addTrail(trail []vec, unique bool) {
	if d.trails[trail[0]] == nil {
		d.trails[trail[0]] = make([][]vec, 0, 100)
	}
	if !slices.ContainsFunc(d.trails[trail[0]], func(t []vec) bool {
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
