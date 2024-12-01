package cmd

import (
	day1 "github.com/lordmoocow/aoc24/internal/day01"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a daily challenge",
	Run: func(cmd *cobra.Command, args []string) {
		// day, err := strconv.Atoi(args[0])
		// if err != nil {
		// 	panic(err)
		// }
		//TODO: actually do this properly

		day1.RunPartOne()
		day1.RunPartTwo()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
