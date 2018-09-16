package word

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/x-color/mkr/test"
)

func TestWordCmdNoOption(t *testing.T) {
	cmd := []string{"word", "testing", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\nworld\n\x1b[32mtesting\x1b[0m now\n"
	actual := test.ExecuteMarkCmd(NewWordCmd, cmd, inputString)
	if actual != expected {
		msg := "Output text was not marked default."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestWordCmdMultiArgs(t *testing.T) {
	cmd := []string{"word", "world", "testing", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[32mworld\x1b[0m\n\x1b[32mtesting\x1b[0m now\n"
	actual := test.ExecuteMarkCmd(NewWordCmd, cmd, inputString)
	if actual != expected {
		msg := "Output text was not marked default."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestWordCmdSetColor(t *testing.T) {
	cmd := []string{"word", "testing", "-b", "yellow", "-c", "red", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\nworld\n\x1b[43m\x1b[31mtesting\x1b[0m now\n"
	actual := test.ExecuteMarkCmd(NewWordCmd, cmd, inputString)
	if actual != expected {
		msg := "Output text was not colored."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestWordCmdSetFormat(t *testing.T) {
	cmd := []string{"word", "testing", "--bold", "--italic", "--hide", "--underline", "--strikethrough", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\nworld\n\x1b[32m\x1b[1m\x1b[3m\x1b[4m\x1b[8m\x1b[9mtesting\x1b[0m now\n"
	actual := test.ExecuteMarkCmd(NewWordCmd, cmd, inputString)
	if actual != expected {
		msg := "Output text was not formated and colored default."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestWordCmdUseCfg(t *testing.T) {
	file, teardown := test.SetupCfgFile(t)
	defer teardown()

	cmd := []string{"word", "testing", "--config", file}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\nworld\n\x1b[33m\x1b[1m\x1b[3mtesting\x1b[0m now\n"
	actual := test.ExecuteMarkCmd(NewWordCmd, cmd, inputString)
	viper.Reset() // for other test code
	if actual != expected {
		msg := "Output text was not marked refering to config file."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}
