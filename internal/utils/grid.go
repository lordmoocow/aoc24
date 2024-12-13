package utils

import "iter"

type Vec struct {
	x, y int
}

func (a *Vec) Add(b *Vec) {
	a.x += b.x
	a.y += b.y
}

type Point Vec

func (a *Point) Add(b *Vec) Point {
	return Point{
		a.x + b.x,
		a.y + b.y,
	}
}

type GridConstraint = comparable

type Grid[T GridConstraint] struct {
	cells []T
	width int
}

func (g *Grid[T]) Bounds() Point {
	return Point{g.width - 1, (len(g.cells) / g.width) - 1}
}

func (g *Grid[T]) InBounds(p Point) bool {
	b := g.Bounds()
	return p.x >= 0 && p.x <= b.x && p.y >= 0 && p.y <= b.y
}

func (g *Grid[T]) Get(p Point) (v T, ok bool) {
	if !g.InBounds(p) {
		return v, false
	}

	return g.cells[g.width*p.y+p.x], true
}

func (g *Grid[T]) Set(p Point, v T) bool {
	if !g.InBounds(p) {
		return false
	}

	g.cells[g.width*p.y+p.x] = v
	return true
}

func (g *Grid[T]) Find(v T) iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for i, cell := range g.cells {
			if cell == v {
				p := Point{i % g.width, i / g.width}
				if !yield(p) {
					return
				}
			}
		}
	}
}

var (
	up    = Vec{0, -1}
	down  = Vec{0, 1}
	left  = Vec{1, 0}
	right = Vec{-1, 0}
)

func (g *Grid[T]) NeighboursUDLR(p Point) iter.Seq2[Point, T] {
	return func(yield func(Point, T) bool) {
		neighbors := []Point{
			p.Add(&up),
			p.Add(&down),
			p.Add(&left),
			p.Add(&right),
		}
		for _, n := range neighbors {
			if v, ok := g.Get(n); ok {
				if !yield(n, v) {
					return
				}
			}
		}
	}
}
