package line

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	mkr "github.com/x-color/mkr/marker"
	"github.com/x-color/mkr/reader"
)

var marker mkr.Marker
var after int
var before int

func preRunLineCmd(cmd *cobra.Command, args []string) {
	viper.Unmarshal(&marker)
}

func runLineCmd(cmd *cobra.Command, args []string) {
	lines := make(chan string, 100)
	go reader.ReadFile(reader.Input, lines)
	// TODO: Check speed these patterns
	//       1. This code
	//       2. Using goroutine this function
	buffer := []string{}
	afterLineNum := after + 1
	for line := range lines {
		if len(buffer) > before {
			cmd.Println(buffer[0])
			buffer = buffer[1:]
		}

		buffer = append(buffer, line)
		for _, pattern := range args {
			if strings.Contains(line, pattern) {
				afterLineNum = 0 // This counted from matched line
				break
			}
		}
		// Print and color matched line and lines before and after it
		if afterLineNum <= after {
			for _, l := range buffer {
				cmd.Println(marker.MarkText(l))
			}
			buffer = buffer[:0]
			afterLineNum++
		}
	}
	// Print lines stacked to print it before matched line in buffer
	for _, l := range buffer {
		cmd.Println(l)
	}
}

// NewLineCmd generates new line command
func NewLineCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "line pattern...",
		Aliases: []string{"l"},
		Short:   "Mark lines containing a text matched pattern",
		Long:    `Line (mrk mark line) marks lines containing a text matched given pattern in standard input.`,
		Example: "  mkr mark line -c yellow --bold message1",
		Args:    cobra.MinimumNArgs(1),
		PreRun:  preRunLineCmd,
		Run:     runLineCmd,
	}

	cmd.Flags().SortFlags = false
	cmd.InheritedFlags().SortFlags = false
	// Marking lines range options
	cmd.Flags().IntVarP(&after,
		"after", "A", 0, "color `num` lines after matched line")
	cmd.Flags().IntVarP(&before,
		"before", "B", 0, "color `num` lines before matched line")

	return cmd
}
