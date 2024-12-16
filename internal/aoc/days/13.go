package days

import (
	"slices"
	"strconv"
	"strings"

	"github.com/lordmoocow/aoc24/internal/utils"
)

type Day13 struct {
	machines []machine
}

type machine struct {
	a, b  utils.Vec
	prize utils.Point
}

func (d *Day13) Init(input string) {
	machines := strings.Split(strings.TrimSpace(input), "\n\n")
	d.machines = make([]machine, len(machines))
	for i, m := range machines {
		lines := strings.Split(m, "\n")
		lineA := strings.Fields(lines[0])
		aX, _ := strconv.Atoi(lineA[2][1 : len(lineA[2])-1])
		aY, _ := strconv.Atoi(lineA[3][1:])
		lineB := strings.Fields(lines[1])
		bX, _ := strconv.Atoi(lineB[2][1 : len(lineB[2])-1])
		bY, _ := strconv.Atoi(lineB[3][1:])
		lineP := strings.Fields(lines[2])
		pX, _ := strconv.Atoi(lineP[1][2 : len(lineP[1])-1])
		pY, _ := strconv.Atoi(lineP[2][2:])
		d.machines[i] = machine{
			a:     utils.Vec{X: aX, Y: aY},
			b:     utils.Vec{X: bX, Y: bY},
			prize: utils.Point{X: pX, Y: pY},
		}
	}
}

type query struct {
	cost int
	d    int
	p    utils.Point
}

func (d *Day13) findPrize(machine machine) int { //([]utils.Point, int) {
	start := utils.Point{X: 0, Y: 0}
	visited := map[utils.Point]bool{}
	visited[start] = true
	//path := make([]utils.Point, 0, 100)

	queue := make([]query, 1, 10)
	queue[0] = query{cost: 0, p: start}
	for {
		if len(queue) == 0 {
			break
		}
		x := queue[0]
		queue = queue[1:]
		//path = append(path, x.p)
		if x.p == machine.prize {
			return x.cost //path, x.cost
		}

		distance := x.p.Dist(&machine.prize).MagSqr()
		for i, n := range []utils.Vec{machine.a, machine.b} {
			nP := x.p.Add(&n)
			// are we getting closer?
			if !visited[nP] {
				nD := nP.Dist(&machine.prize).MagSqr()
				if nD < distance {
					token := 3
					if i == 1 {
						token = 1
					}
					visited[x.p] = true
					queue = append(queue, query{cost: x.cost + token, d: nD, p: nP})
					slices.SortFunc(queue, func(a, b query) int { return (a.d + a.cost) - (b.d + b.cost) })
				}
			}
		}
	}
	return 0
}

func (d *Day13) PartOne() (result int) {
	for _, m := range d.machines {
		tokens := d.findPrize(m)
		result += tokens
	}
	return result
}

func (d *Day13) PartTwo() (result int) {

	return result
}
