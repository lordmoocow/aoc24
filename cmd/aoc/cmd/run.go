package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/lordmoocow/aoc24/internal/aoc"
	"github.com/spf13/cobra"
)

var (
	parts []int
)

var runCmd = &cobra.Command{
	Use:   "run <day> [--part n]... <input_file>",
	Short: "Run an advent of code puzzle",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		day, input, err := parseArgs(args)
		if err != nil {
			return err
		}

		fmt.Printf("Day %d\n", day)
		for _, part := range parts {
			result, err := aoc.Run(day, part, input)
			if err != nil {
				return err
			}
			fmt.Printf("  Part %d: %d\n", part, result)
		}

		return nil
	},
}

func parseArgs(args []string) (day int, input string, err error) {
	day, err = strconv.Atoi(args[0])
	if err != nil {
		return day, input, err
	}
	if day > 25 || day < 1 {
		return day, input, errors.New("day out of range (1-25)")
	}

	file, err := os.Open(args[1])
	if err != nil {
		return day, input, err
	}
	buf, _ := io.ReadAll(file)
	input = string(buf)

	return day, input, err
}

func init() {
	runCmd.Flags().IntSliceVarP(&parts, "part", "p", []int{1, 2}, "the part of the challenge to run")
	rootCmd.AddCommand(runCmd)
}
