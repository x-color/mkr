package number

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/x-color/mkr/test"
)

func TestNumberCmdOneLine(t *testing.T) {
	cmd := []string{"number", "2", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[32mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewNumberCmd, cmd, inputString)
	if actual != expected {
		msg := "Specified line is not marked."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestNumberCmdEvenLines(t *testing.T) {
	cmd := []string{"number", "even", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[32mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewNumberCmd, cmd, inputString)
	if actual != expected {
		msg := "Even lines is not marked."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestNumberCmdMultipleLines(t *testing.T) {
	cmd := []string{"number", "2n", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[32mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewNumberCmd, cmd, inputString)
	if actual != expected {
		msg := "2n lines is not marked."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestNumberCmdBeforeLines(t *testing.T) {
	cmd := []string{"number", "~2", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "\x1b[32mhello\x1b[0m\n\x1b[32mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewNumberCmd, cmd, inputString)
	if actual != expected {
		msg := "Specified line and lines before it are not marked."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestNumberCmdAfterLines(t *testing.T) {
	cmd := []string{"number", "2~", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[32mworld\x1b[0m\n\x1b[32mtesting now\x1b[0m\n"
	actual := test.ExecuteMarkCmd(NewNumberCmd, cmd, inputString)
	if actual != expected {
		msg := "Specified line and lines after it are not marked."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestNumberCmdLinesInRange(t *testing.T) {
	cmd := []string{"number", "1~3", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "\x1b[32mhello\x1b[0m\n\x1b[32mworld\x1b[0m\n\x1b[32mtesting now\x1b[0m\n"
	actual := test.ExecuteMarkCmd(NewNumberCmd, cmd, inputString)
	if actual != expected {
		msg := "Lines in specified range are not marked."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestNumberCmdMultiArgs(t *testing.T) {
	cmd := []string{"number", "1", "3", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "\x1b[32mhello\x1b[0m\nworld\n\x1b[32mtesting now\x1b[0m\n"
	actual := test.ExecuteMarkCmd(NewNumberCmd, cmd, inputString)
	if actual != expected {
		msg := "Specified lines are not marked."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestNumberCmdSetColor(t *testing.T) {
	cmd := []string{"number", "2", "-b", "yellow", "-c", "red", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[43m\x1b[31mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewNumberCmd, cmd, inputString)
	if actual != expected {
		msg := "Output text is not marked specified color."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestNumberCmdSetFormat(t *testing.T) {
	cmd := []string{"number", "2", "--bold", "--italic", "--hide", "--underline", "--strikethrough", "--config", "no-config-file.dummy.yaml"}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[32m\x1b[1m\x1b[3m\x1b[4m\x1b[8m\x1b[9mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewNumberCmd, cmd, inputString)
	if actual != expected {
		msg := "Output text is not marked specified format."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}

func TestNumberCmdUseCfg(t *testing.T) {
	file, teardown := test.SetupCfgFile(t)
	defer teardown()

	cmd := []string{"number", "2", "--config", file}
	inputString := "hello\nworld\ntesting now"

	expected := "hello\n\x1b[33m\x1b[1m\x1b[3mworld\x1b[0m\ntesting now\n"
	actual := test.ExecuteMarkCmd(NewNumberCmd, cmd, inputString)
	viper.Reset() // for other test code
	if actual != expected {
		msg := "Output text was not marked refering to config file."
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
	}
}
