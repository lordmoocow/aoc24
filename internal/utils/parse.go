package utils

import (
	"strconv"
	"strings"
)

type ParseOptions struct {
	row string
	col string
}

var DefaultParseOptions = ParseOptions{
	row: "\n",
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

	buf := make([]rune, 0, 10)
	for i, v := range strings.TrimSpace(input) {
		buf = append(buf, v)

		// check if we have reached the end of a row
		if len(options.row) > 0 {
			if len(buf) >= len(options.row) && string(buf[len(buf)-len(options.row):]) == options.row {
				if g.width == 0 {
					g.width = i
				}
				if len(buf) > len(options.row) {
					g.cells = append(g.cells, val(buf[:len(buf)-len(options.row)]))
				}
				buf = buf[:0]
				continue
			}
		}

		// check if we have a complete cell value
		if len(options.col) == 0 || (len(buf) >= len(options.col) && string(buf[len(buf)-len(options.col):]) == options.col) {
			g.cells = append(g.cells, val(buf[:len(buf)-len(options.col)]))
			buf = buf[:0]
			continue
		}
	}

	if len(buf) > 0 {
		g.cells = append(g.cells, val(buf))
	}

	return g
}
