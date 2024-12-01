package main

import (
	"time"

	"github.com/lordmoocow/aoc24/cmd/aoc/cmd"
)

var (
	version = "dev"
	commit  = ""
	date    = time.Now().Format(time.RFC3339)
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
