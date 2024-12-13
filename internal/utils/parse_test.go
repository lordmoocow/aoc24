package utils

import "testing"

func TestParseGrid(t *testing.T) {
	tests := map[string]struct {
		input      string
		options    ParseOptions
		cols, rows int
	}{
		"5x5 no space between cols": {
			input: `12345
12345
12345
12345
12345`,
			options: ParseOptions{col: ""},
			cols:    5,
			rows:    5,
		},
		"5x5 space between cols": {
			input: `1 2 3 4 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5`,
			options: ParseOptions{col: " "},
			cols:    5,
			rows:    5,
		},
		"1x10 space between cols": {
			input:   `1 2 3 4 5 1 2 3 4 5`,
			options: ParseOptions{col: " "},
			cols:    10,
			rows:    1,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := ParseGrid(test.input, test.options, func(v []rune) string { return string(v) })
			expectedCells := test.cols * test.rows
			if len(got.cells) != expectedCells {
				t.Fatalf("ParseGrid(%q) returned %q; expected %q cells", test.input, got.cells, expectedCells)
			}
		})
	}
}
