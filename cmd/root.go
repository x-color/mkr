package cmd

import (
	"os"

	"github.com/spf13/cobra"
	init_ "github.com/x-color/mkr/cmd/init"
	mark "github.com/x-color/mkr/cmd/mark"
	"github.com/x-color/mkr/reader"
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "mkr",
		Long: "Mkr is marking tool. It marks a text in standard input.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(mark.NewMarkCmd())
	cmd.AddCommand(init_.NewInitCmd())

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := rootCmd()
	cmd.SetOutput(os.Stdout)
	reader.Input = os.Stdin
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}
