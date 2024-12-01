package day1

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var input string

func RunPartOne() {
	a := []int{}
	b := []int{}
	for _, v := range strings.Split(input, "\n") {
		x := strings.Fields(v)
		if len(x) > 0 {
			x1, _ := strconv.Atoi(x[0])
			a = append(a, x1)
			x2, _ := strconv.Atoi(x[1])
			b = append(b, x2)
		}
	}

	slices.Sort(a)
	slices.Sort(b)

	result := 0
	for i := 0; i < len(a); i++ {
		distance := a[i] - b[i]
		if distance < 0 {
			distance *= -1
		}
		result += distance
	}

	fmt.Println(result)
}
