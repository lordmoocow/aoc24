package utils

import (
	"bufio"
	"strconv"
	"strings"
)

type ParseOptions struct {
	col string
}

var DefaultParseOptions = ParseOptions{
	col: "",
}

func ParseIntGrid(input string, options ParseOptions) Grid[int] {
	return ParseGrid(input, options, func(v []rune) int {
		i, _ := strconv.Atoi(string(v))
		return i
	})
}

func ParseRuneGrid(input string, options ParseOptions) Grid[rune] {
	return ParseGrid(input, options, func(v []rune) rune {
		return rune(v[0])
	})
}

func ParseGrid[T GridConstraint](input string, options ParseOptions, val func(v []rune) T) Grid[T] {
	g := Grid[T]{}

	rowscan := bufio.NewScanner(strings.NewReader(input))
	for {
		if !rowscan.Scan() {
			break
		}

		cols := strings.Split(rowscan.Text(), options.col)
		if g.width == 0 {
			g.width = len(cols)
		}
		for _, v := range cols {
			g.cells = append(g.cells, val([]rune(v)))
		}
	}

	return g
}
