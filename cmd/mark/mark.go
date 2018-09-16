package mark

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-color/mkr/cmd/mark/line"
	"github.com/x-color/mkr/cmd/mark/number"
	"github.com/x-color/mkr/cmd/mark/word"
)

var cfgFile string

// NewMarkCmd generates new mark command
func NewMarkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark",
		Short: "Mark a text",
		Long: `Mark (mrk mark) marks a text by using subcommands.

Use -b, -c to specify marking color. Settable color is
 - name   'none','black','blue','cyan','green','magenta','red','yellow'
 - number 0~255, it is supported by some terminals`,
		Example: "  mkr mark line message1",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.PersistentFlags().SortFlags = false
	cmd.Flags().SortFlags = false

	// Marking colors options
	cmd.PersistentFlags().StringP("background", "b", "none", "`color` background")
	cmd.PersistentFlags().StringP("charactor", "c", "green", "`color` charactor")
	// Marking formats options
	cmd.PersistentFlags().Bool("bold", false, "make text bold")
	cmd.PersistentFlags().Bool("hide", false, "hide text")
	cmd.PersistentFlags().Bool("italic", false, "make text italic")
	cmd.PersistentFlags().Bool("strikethrough", false, "strikethrough text")
	cmd.PersistentFlags().Bool("underline", false, "underline text")
	// Config file
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config `file` (default is $HOME/.mkr.yaml)")

	viper.BindPFlag("color.background", cmd.PersistentFlags().Lookup("background"))
	viper.BindPFlag("color.charactor", cmd.PersistentFlags().Lookup("charactor"))
	viper.BindPFlag("format.bold", cmd.PersistentFlags().Lookup("bold"))
	viper.BindPFlag("format.hide", cmd.PersistentFlags().Lookup("hide"))
	viper.BindPFlag("format.italic", cmd.PersistentFlags().Lookup("italic"))
	viper.BindPFlag("format.strikethrough", cmd.PersistentFlags().Lookup("strikethrough"))
	viper.BindPFlag("format.underline", cmd.PersistentFlags().Lookup("underline"))

	cmd.AddCommand(line.NewLineCmd())
	cmd.AddCommand(number.NewNumberCmd())
	cmd.AddCommand(word.NewWordCmd())

	return cmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".mkr")
	}
	viper.ReadInConfig()
}
