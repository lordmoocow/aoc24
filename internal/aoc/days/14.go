package days

import (
	"math"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/lordmoocow/aoc24/internal/utils"
)

type Day14 struct {
	robots []robot
	room   utils.Vec
}

type robot struct {
	pos utils.Point
	vel utils.Vec
}

func (d *Day14) Init(input string) {
	d.room = utils.Vec{X: 100, Y: 102}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	d.robots = make([]robot, len(lines))
	for i, line := range lines {
		split := strings.Index(line, " ")
		p := strings.Split(line[2:split], ",")
		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])
		v := strings.Split(line[split+3:], ",")
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])
		d.robots[i] = robot{
			pos: utils.Point{X: px, Y: py},
			vel: utils.Vec{X: vx, Y: vy},
		}
	}
}

func (d *Day14) simulate(steps int) {
	for i, r := range d.robots {
		v := r.vel.Scale(steps)
		p := r.pos.Add(&v)

		p.X %= d.room.X + 1
		p.Y %= d.room.Y + 1

		if p.X < 0 {
			p.X += d.room.X + 1
		}
		if p.Y < 0 {
			p.Y += d.room.Y + 1
		}

		r.pos = p
		d.robots[i] = r
	}
}

func (d *Day14) safetyFactor() int {
	mid := utils.Point{
		X: d.room.X / 2,
		Y: d.room.Y / 2,
	}
	quads := [4]int{}
	for _, r := range d.robots {
		if r.pos.X < mid.X {
			if r.pos.Y < mid.Y {
				quads[0]++
			} else if r.pos.Y > mid.Y {
				quads[2]++
			}
		} else if r.pos.X > mid.X {
			if r.pos.Y < mid.Y {
				quads[1]++
			} else if r.pos.Y > mid.Y {
				quads[3]++
			}
		}
	}
	return quads[0] * quads[1] * quads[2] * quads[3]
}

func (d *Day14) distribution() (int, int) {
	meanX := 0.0
	meanY := 0.0
	for _, r := range d.robots {
		meanX += float64(r.pos.X)
		meanY += float64(r.pos.Y)
	}
	meanX /= float64(len(d.robots))
	meanY /= float64(len(d.robots))

	varX := 0.0
	varY := 0.0
	for _, r := range d.robots {
		varX += math.Pow(float64(r.pos.X)-meanX, 2)
		varY += math.Pow(float64(r.pos.Y)-meanY, 2)
	}
	varX /= float64(len(d.robots))
	varY /= float64(len(d.robots))

	sdX := math.Sqrt(varX)
	sdY := math.Sqrt(varY)

	return int(sdX), int(sdY)
}

func (d *Day14) PartOne() (result int) {
	d.simulate(100)
	return d.safetyFactor()
}

func (d *Day14) PartTwo() (result int) {
	screen, _ = tcell.NewScreen()
	if screen != nil {
		screen.Init()
		screen.Clear()
		screen.Resize(d.room.X, d.room.Y, 0, 0)
		defer func() {
			screen.Fini()
		}()
	}

screenLoop:
	for {
		for {
			result += 1
			d.simulate(1)
			dx, dy := d.distribution()
			if dx+dy < 40 {
				break
			}
		}

		if screen != nil {
			d.render()
			screen.Show()

			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					break screenLoop
				}
			}
		}
	}
	return result
}

var screen tcell.Screen

func (d *Day14) render() {
	for y := range d.room.Y {
		for x := range d.room.X {
			screen.SetContent(x, y, '.', nil, tcell.StyleDefault)
		}
	}
	for _, r := range d.robots {
		count := '1'
		c, _, _, _ := screen.GetContent(r.pos.X, r.pos.Y)
		if c != '.' {
			i, _ := strconv.Atoi(string(c))
			i++
			count = rune(strconv.Itoa(i)[0])
		}
		screen.SetContent(r.pos.X, r.pos.Y, count, nil, tcell.StyleDefault)
	}
}
