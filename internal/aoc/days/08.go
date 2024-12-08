package days

import (
	"strings"
)

type Day8 struct {
	city       map[vec]bool
	antennas   map[antenna][]vec
	maxX, maxY int
}

type antenna struct {
	freq rune
}

type antinode struct{}

func (d *Day8) Init(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	d.city = map[vec]bool{}
	d.antennas = map[antenna][]vec{}
	for y, row := range lines {
		d.maxY = y
		for x, col := range row {
			d.maxX = x
			if col != '.' {
				a := antenna{col}
				pos := vec{x, y}
				d.antennas[a] = append(d.antennas[a], pos)
			}
		}
	}
}

func (d *Day8) PartOne() (result int) {
	d.locateAntinodes()
	for _, hasAntinode := range d.city {
		if hasAntinode {
			result++
		}
	}
	return result
}

func (d *Day8) locateAntinodes() {
	for _, array := range d.antennas {
		// single antenna cannot produce antinodes
		if len(array) == 1 {
			continue
		}

		// for each antenna of this freq
		for _, a := range array {
			// look at location of all _other_ matching freq antennas
			for _, b := range array {
				// skip self
				if a == b {
					continue
				}

				// by definition two points always make a straight line
				// so just need to check if the distance x 2 is within bounds
				dx := b.x - a.x
				dy := b.y - a.y
				nodeLoc := vec{b.x + dx, b.y + dy}
				if nodeLoc.x >= 0 && nodeLoc.y >= 0 && nodeLoc.x <= d.maxX && nodeLoc.y <= d.maxY {
					d.city[nodeLoc] = true
				}
			}
		}
	}
}

func (d *Day8) PartTwo() (result int) {
	return result
}
