package word

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	mkr "github.com/x-color/mkr/marker"
	"github.com/x-color/mkr/reader"
)

var marker mkr.Marker

func preRunWordCmd(cmd *cobra.Command, args []string) {
	viper.Unmarshal(&marker)
}

func runWordCmd(cmd *cobra.Command, args []string) {
	lines := make(chan string, 100)
	go reader.ReadFile(reader.Input, lines)
	// TODO: Check speed these patterns
	//       1. This code
	//       2. Using goroutine this function
	for line := range lines {
		for _, pattern := range args {
			if strings.Contains(line, pattern) {
				line = strings.Replace(line, pattern, marker.MarkText(pattern), -1)
			}
		}
		cmd.Println(line)
	}
}

// NewWordCmd generates new word command
func NewWordCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "word pattern...",
		Aliases: []string{"w"},
		Short:   "Mark a text matched pattern",
		Long:    `Word (mrk mark word) marks a text matched given pattern in standard input.`,
		Example: "  mkr mark word -c yellow --bold message1",
		Args:    cobra.MinimumNArgs(1),
		PreRun:  preRunWordCmd,
		Run:     runWordCmd,
	}

	cmd.InheritedFlags().SortFlags = false

	return cmd
}
