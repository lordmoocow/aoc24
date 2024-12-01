package cmd

import (
	day1 "github.com/lordmoocow/aoc24/internal/day01"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a daily challenge",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
