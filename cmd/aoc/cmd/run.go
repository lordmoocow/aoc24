package cmd

import (
	"fmt"
	"strconv"

	"github.com/lordmoocow/aoc24/assets"
	"github.com/lordmoocow/aoc24/internal/aoc"
	day1 "github.com/lordmoocow/aoc24/internal/day01"
	day2 "github.com/lordmoocow/aoc24/internal/day02"
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
		return &day1.Day1{}
	case 2:
		return &day2.Day2{}
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
