package days

import (
	"math"
	"strconv"
	"strings"
)

type Day11 struct {
	stones []int
}

func (d *Day11) Init(input string) {
	stones := strings.Fields(strings.TrimSpace(input))
	d.stones = make([]int, len(stones))
	for i, stone := range stones {
		s, _ := strconv.Atoi(stone)
		d.stones[i] = s
	}
}

func (d *Day11) blink() {
	stones := make([]int, 0, len(d.stones)+len(d.stones)/25)
	for _, stone := range d.stones {
		if stone == 0 {
			stones = append(stones, 1)
		} else if int(math.Log10(float64(stone)))%2 == 1 {
			s := strconv.Itoa(stone)
			mid := len(s) / 2
			stone, _ = strconv.Atoi(s[:mid])
			stones = append(stones, stone)
			stone, _ = strconv.Atoi(s[mid:])
			stones = append(stones, stone)
		} else {
			stones = append(stones, stone*2024)
		}
	}
	d.stones = stones
}

func (d *Day11) PartOne() (result int) {
	for range 25 {
		d.blink()
	}
	return len(d.stones)
}

func (d *Day11) PartTwo() (result int) {
	return result
}
