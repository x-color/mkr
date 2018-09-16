package line

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/x-color/mkr/test"
)

func TestLineCmdNoOption(t *testing.T) {
	cmd := []string{"line", "world", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[32mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewLineCmd, cmd, inputString)
	if actual != expected {
		msg := "Output text was not marked default."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestLineCmdMultiArgs(t *testing.T) {
	cmd := []string{"line", "world", "now", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[32mworld\x1b[0m\n\x1b[32mtesting now\x1b[0m\n"
	actual := test.ExecuteMarkCmd(NewLineCmd, cmd, inputString)
	if actual != expected {
		msg := "Output text was not marked default."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestLineCmdSetColor(t *testing.T) {
	cmd := []string{"line", "world", "-b", "yellow", "-c", "red", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[43m\x1b[31mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewLineCmd, cmd, inputString)
	if actual != expected {
		msg := "Output text was not colored."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestLineCmdSetFormat(t *testing.T) {
	cmd := []string{"line", "world", "--bold", "--italic", "--hide", "--underline", "--strikethrough", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[32m\x1b[1m\x1b[3m\x1b[4m\x1b[8m\x1b[9mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewLineCmd, cmd, inputString)
	if actual != expected {
		msg := "Output text was not formated and colored default."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestLineCmdMarkAfter(t *testing.T) {
	cmd := []string{"line", "world", "-A", "1", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[32mworld\x1b[0m\n\x1b[32mtesting now\x1b[0m\n"
	actual := test.ExecuteMarkCmd(NewLineCmd, cmd, inputString)
	if actual != expected {
		msg := "Matched and after text were not colored default."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestLineCmdMarkBefore(t *testing.T) {
	cmd := []string{"line", "world", "-B", "1", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "\x1b[32mhello\x1b[0m\n\x1b[32mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewLineCmd, cmd, inputString)
	if actual != expected {
		msg := "Matched and after text were not colored default."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestLineCmdUseCfg(t *testing.T) {
	file, teardown := test.SetupCfgFile(t)
	defer teardown()

	cmd := []string{"line", "world", "--config", file}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[33m\x1b[1m\x1b[3mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewLineCmd, cmd, inputString)
	viper.Reset() // for other test code
	if actual != expected {
		msg := "Output text was not marked refering to config file."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}
