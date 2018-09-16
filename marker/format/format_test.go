package format

import (
	"testing"
)

func TestGenStyleCode(t *testing.T) {
	testCases := []struct {
		format Format
		output string
	}{
		{
			format: Format{
				Bold:          true,
				Italic:        true,
				Underline:     true,
				Hide:          true,
				Strikethrough: true,
			},
			output: "\x1b[1m\x1b[3m\x1b[4m\x1b[8m\x1b[9m",
		},
		{
			format: Format{
				Bold:          false,
				Italic:        false,
				Underline:     false,
				Hide:          false,
				Strikethrough: false,
			},
			output: "",
		},
	}

	for _, tc := range testCases {
		expected := tc.output
		actual := tc.format.GenFormatCode()
		if actual != expected {
			msg := "Did not generate correct format code.\n"
			t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
		}
	}
}
