package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/lordmoocow/aoc24/internal/aoc"
	"github.com/spf13/cobra"
)

var (
	day   int
	parts []int
)

var runCmd = &cobra.Command{
	Use:   "run -d day [-p part]... INPUT",
	Short: "Run an advent of code puzzle",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := parseArgs(args)
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

func parseArgs(args []string) (input string, err error) {
	file, err := os.Open(args[0])
	if err != nil {
		return input, err
	}
	buf, _ := io.ReadAll(file)
	input = string(buf)

	return input, err
}

func init() {
	runCmd.Flags().IntVarP(&day, "day", "d", 0, "the day of the calendar to run")
	runCmd.MarkFlagRequired("day")
	runCmd.Flags().IntSliceVarP(&parts, "part", "p", []int{1, 2}, "the part of the challenge to run")
	rootCmd.AddCommand(runCmd)
}
