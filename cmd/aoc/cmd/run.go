package cmd

import (
	"fmt"
	"strconv"

	"github.com/lordmoocow/aoc24/internal/aoc"
	"github.com/lordmoocow/aoc24/internal/assets"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a daily challenge",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dayNumber, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid day")
		}

		partNumber, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("invalid part")
		}

		input := assets.InputData(dayNumber)
		result, err := aoc.Run(dayNumber, partNumber, input)
		if err != nil {
			fmt.Printf("Day %d, Part %d: %s\n", dayNumber, partNumber, err.Error())
			return
		}
		fmt.Printf("Day %d, Part %d: %d\n", dayNumber, partNumber, result)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
