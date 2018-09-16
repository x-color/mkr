package number

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	mkr "github.com/x-color/mkr/marker"
	"github.com/x-color/mkr/reader"
)

var marker mkr.Marker
var checkers []func(int) bool

func preRunNumberCmd(cmd *cobra.Command, args []string) {
	viper.Unmarshal(&marker)
	checkers = parseRanges(args)
}

func runNumberCmd(cmd *cobra.Command, args []string) {
	lines := make(chan string, 100)
	go reader.ReadFile(reader.Input, lines)
	// TODO: Check speed these patterns
	//       1. This code
	//       2. Using goroutine this function
	lineNum := 0
	for line := range lines {
		lineNum++
		if inRanges(lineNum, checkers) {
			line = marker.MarkText(line)
		}
		cmd.Println(line)
	}
}

// NewNumberCmd generates new number command
func NewNumberCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "number range...",
		Aliases: []string{"n"},
		Short:   "Mark lines in a range",
		Long: `Number (mrk mark number) marks lines in a given range in standard input.

range is an integer, a type of number, a range of integers, or multiple of integer.
 - N    line N counted from 1
 - N~   from line N to the end of the file
 - ~M   from the first of the file to line M
 - N~M  from line N to line M
 - Nn   line multiples of N
 - even even-numbered lines
 - odd  odd-numbered lines`,
		Example: "  mkr mark number -c yellow --bold 3~9",
		Args:    cobra.MinimumNArgs(1),
		PreRun:  preRunNumberCmd,
		Run:     runNumberCmd,
	}

	cmd.InheritedFlags().SortFlags = false

	return cmd
}
