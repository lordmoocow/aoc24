package cmd

import (
	"fmt"
	"strconv"

	"github.com/lordmoocow/aoc24/internal/aoc"
	"github.com/lordmoocow/aoc24/internal/aoc/days"
	"github.com/lordmoocow/aoc24/internal/assets"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a daily challenge",
	Run: func(cmd *cobra.Command, args []string) {
		dayNumber, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		partNumber, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}

		input := assets.InputData(dayNumber)
		result := run(dayNumber, partNumber, input)
		fmt.Printf("Day %d, Part %d: %d\r\n", dayNumber, partNumber, result)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func getDay(day int) aoc.Day {
	switch day {
	case 1:
		return &days.Day1{}
	case 2:
		return &days.Day2{}
	case 3:
		return &days.Day3{}
	}

	return nil
}

func run(day, part int, input string) int {
	d := getDay(day)
	d.Init(input)

	switch part {
	case 1:
		return d.PartOne()
	case 2:
		return d.PartTwo()
	}
	return -1
}
