package test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-color/mkr/reader"
)

var cfgFile string

func dummyRootCmd(newCmd func() *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use: "mkr",
	}
	cmd.AddCommand(newCmd())
	return cmd
}

func dummyMarkCmd(newCmd func() *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use: "mark",
	}

	cmd.PersistentFlags().StringP("background", "b", "none", "")
	cmd.PersistentFlags().StringP("charactor", "c", "green", "")
	cmd.PersistentFlags().Bool("bold", false, "")
	cmd.PersistentFlags().Bool("hide", false, "")
	cmd.PersistentFlags().Bool("italic", false, "")
	cmd.PersistentFlags().Bool("strikethrough", false, "")
	cmd.PersistentFlags().Bool("underline", false, "")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	viper.BindPFlag("color.background", cmd.PersistentFlags().Lookup("background"))
	viper.BindPFlag("color.charactor", cmd.PersistentFlags().Lookup("charactor"))
	viper.BindPFlag("format.bold", cmd.PersistentFlags().Lookup("bold"))
	viper.BindPFlag("format.hide", cmd.PersistentFlags().Lookup("hide"))
	viper.BindPFlag("format.italic", cmd.PersistentFlags().Lookup("italic"))
	viper.BindPFlag("format.strikethrough", cmd.PersistentFlags().Lookup("strikethrough"))
	viper.BindPFlag("format.underline", cmd.PersistentFlags().Lookup("underline"))

	cmd.AddCommand(newCmd())
	return cmd
}

// SetupCfgFile sets up tmp config file
func SetupCfgFile(t *testing.T) (string, func()) {
	t.Helper()
	tmpFile, err := ioutil.TempFile("", "tmptest")
	if err != nil {
		t.Fatalf("Error occurred in setup process before testing.\n%s\n", err)
	}
	data := `color:
    background: none
    charactor: yellow
format:
    bold: true
    italic: true
    hide: false
    underline: false
    strikethrough: false`
	if err := ioutil.WriteFile(tmpFile.Name(), []byte(data), 0644); err != nil {
		t.Fatalf("Error occurred in setup process before testing.\n%s\n", err)
	}
	tmpCfgFile := tmpFile.Name() + ".yaml"
	if err := os.Rename(tmpFile.Name(), tmpCfgFile); err != nil {
		t.Fatalf("Error occurred in setup process before testing.\n%s\n", err)
	}
	return tmpCfgFile, func() { os.Remove(tmpCfgFile) }
}

// ExecuteMarkCmd executes dummy mark command
func ExecuteMarkCmd(newCmd func() *cobra.Command, args []string, inputString string) string {
	cmd := dummyMarkCmd(newCmd)
	reader.Input = bytes.NewBufferString(inputString)
	buf := new(bytes.Buffer)
	cmd.SetOutput(buf)
	cmd.SetArgs(args)
	cmd.Execute()
	return buf.String()
}

// ExecuteRootCmd executes dummy root command
func ExecuteRootCmd(newCmd func() *cobra.Command, args []string, inputString string) string {
	cmd := dummyRootCmd(newCmd)
	reader.Input = bytes.NewBufferString(inputString)
	buf := new(bytes.Buffer)
	cmd.SetOutput(buf)
	cmd.SetArgs(args)
	cmd.Execute()
	return buf.String()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.ReadInConfig()
}
