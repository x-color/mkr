package init

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	mkr "github.com/x-color/mkr/marker"
)

var marker mkr.Marker

func runInitCmd(cmd *cobra.Command, args []string) error {
	viper.Set("color.background", marker.Color.Background)
	viper.Set("color.charactor", marker.Color.Charactor)
	viper.Set("format.bold", marker.Format.Bold)
	viper.Set("format.hide", marker.Format.Hide)
	viper.Set("format.italic", marker.Format.Italic)
	viper.Set("format.strikethrough", marker.Format.Strikethrough)
	viper.Set("format.underline", marker.Format.Underline)

	file := ""
	if len(args) == 0 {
		home, err := homedir.Dir()
		if err != nil {
			return err
		}
		file = home + "/.mkr.yaml"
	} else {
		file = args[0]
	}

	if err := viper.WriteConfigAs(file); err != nil {
		return err
	}
	cmd.Printf("Create new config file '%s'\n", file)
	return nil
}

// NewInitCmd generates new init command
func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [new file]",
		Short: "Create new config file for marking",
		Long: `Init (mrk init) creates new config file (default: $HOME/.mkr.yaml) for default marking style.

Use -b, -c to specify marking color. Settable color is
 - name   'none','black','blue','cyan','green','magenta','red','yellow'
 - number 0~255, it is supported by some terminals`,
		Example: "  mkr init",
		Args:    cobra.MaximumNArgs(1),
		RunE:    runInitCmd,
	}

	cmd.Flags().SortFlags = false
	// Marking colors options
	cmd.Flags().StringVarP(&marker.Color.Background,
		"background", "b", "none", "set default background `color`")
	cmd.Flags().StringVarP(&marker.Color.Charactor,
		"charactor", "c", "green", "set default charactor `color`")
	// Marking formats options
	cmd.Flags().BoolVar(&marker.Format.Bold,
		"bold", false, "set bold format")
	cmd.Flags().BoolVar(&marker.Format.Hide,
		"hide", false, "set hidden format")
	cmd.Flags().BoolVar(&marker.Format.Italic,
		"italic", false, "set italic format")
	cmd.Flags().BoolVar(&marker.Format.Strikethrough,
		"strikethrough", false, "set strikethrough format")
	cmd.Flags().BoolVar(&marker.Format.Underline,
		"underline", false, "set underline format")

	return cmd
}
