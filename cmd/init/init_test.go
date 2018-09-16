package init

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/x-color/mkr/test"
)

func testCheckCfgFile(t *testing.T, file, expected string) {
	t.Helper()
	data, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("Error occurred before checking created file.\n%s\n", err)
	}
	actual := string(data)
	if actual != expected {
		msg := "Created config file is not correctory."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestInitCmdNoOption(t *testing.T) {
	tmpFile := "tmpConfigFile.yaml"
	defer os.Remove(tmpFile)
	defer viper.Reset()

	cmd := []string{"init", tmpFile}
	expected := fmt.Sprintf("Create new config file '%s'\n", tmpFile)
	actual := test.ExecuteRootCmd(NewInitCmd, cmd, "")
	if actual != expected {
		msg := "Output text was not marked default."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}

	expected = `color:
  background: none
  charactor: green
format:
  bold: false
  hide: false
  italic: false
  strikethrough: false
  underline: false
`
	testCheckCfgFile(t, tmpFile, expected)
}

func TestInitCmdSetColor(t *testing.T) {
	tmpFile := "tmpConfigFile.yaml"
	defer os.Remove(tmpFile)
	defer viper.Reset()

	cmd := []string{"init", tmpFile, "-b", "yellow", "-c", "red"}
	expected := fmt.Sprintf("Create new config file '%s'\n", tmpFile)
	actual := test.ExecuteRootCmd(NewInitCmd, cmd, "")
	if actual != expected {
		msg := "Output text was not correctory."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}

	expected = `color:
  background: yellow
  charactor: red
format:
  bold: false
  hide: false
  italic: false
  strikethrough: false
  underline: false
`
	testCheckCfgFile(t, tmpFile, expected)
}

func TestInitCmdSetFormat(t *testing.T) {
	tmpFile := "tmpConfigFile.yaml"
	defer os.Remove(tmpFile)
	defer viper.Reset()

	cmd := []string{"init", tmpFile, "--bold", "--italic", "--hide", "--underline", "--strikethrough"}
	expected := fmt.Sprintf("Create new config file '%s'\n", tmpFile)
	actual := test.ExecuteRootCmd(NewInitCmd, cmd, "")
	if actual != expected {
		msg := "Output text was not formated and colored default."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}

	expected = `color:
  background: none
  charactor: green
format:
  bold: true
  hide: true
  italic: true
  strikethrough: true
  underline: true
`
	testCheckCfgFile(t, tmpFile, expected)
}
