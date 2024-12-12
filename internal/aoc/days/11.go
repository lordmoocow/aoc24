package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day11 struct {
	stones     []int
	evolutions map[string]int
}

func (d *Day11) Init(input string) {
	stones := strings.Fields(strings.TrimSpace(input))
	d.stones = make([]int, len(stones))
	for i, stone := range stones {
		s, _ := strconv.Atoi(stone)
		d.stones[i] = s
	}
}

func (d *Day11) blink(count int) int {
	result := 0
	d.evolutions = map[string]int{}

	for _, stone := range d.stones {
		result += d.evolve(stone, count)
	}
	return result
}

func (d *Day11) evolve(stone, evolutions int) int {
	if evolutions == 0 {
		return 1
	}

	if c, ok := d.evolutions[fmt.Sprintf("%d|%d", stone, evolutions)]; ok {
		return c
	}

	result := 0
	if stone == 0 {
		result += d.evolve(1, evolutions-1)
	} else if int(math.Log10(float64(stone)))%2 == 1 {
		s := strconv.Itoa(stone)
		mid := len(s) / 2
		a, _ := strconv.Atoi(s[:mid])
		b, _ := strconv.Atoi(s[mid:])
		result += d.evolve(a, evolutions-1)
		result += d.evolve(b, evolutions-1)
	} else {
		result += d.evolve(stone*2024, evolutions-1)
	}
	d.evolutions[fmt.Sprintf("%d|%d", stone, evolutions)] = result
	return result
}

func (d *Day11) PartOne() (result int) {
	result = d.blink(25)
	return result
}

func (d *Day11) PartTwo() (result int) {
	result = d.blink(75)
	return result
}
