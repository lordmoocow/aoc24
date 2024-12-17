package utils

import (
	"cmp"
	"slices"
)

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Push(item T) {
	q.data = append(q.data, item)
}

func (q *Queue[T]) Next() T {
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

type pqitem[T any] struct {
	p    int
	item T
}

type PriorityQueue[T any] struct {
	Queue[pqitem[T]]
}

func NewPQ[T any](cap int) PriorityQueue[T] {
	return PriorityQueue[T]{
		Queue: Queue[pqitem[T]]{
			data: make([]pqitem[T], 0, cap),
		},
	}
}

func (q *PriorityQueue[T]) Push(priority int, item T) {
	i, _ := slices.BinarySearchFunc(q.data, priority, func(a pqitem[T], b int) int {
		return cmp.Compare(a.p, b)
	})
	q.data = append(q.data, pqitem[T]{})
	copy(q.data[i+1:], q.data[i:])
	q.data[i] = pqitem[T]{priority, item}
}

func (q *PriorityQueue[T]) Next() (int, T) {
	n := q.Queue.Next()
	return n.p, n.item
}
