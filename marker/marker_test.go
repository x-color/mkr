package marker

import (
	"testing"

	"github.com/x-color/mkr/marker/color"
	"github.com/x-color/mkr/marker/format"
)

func TestGenBackColor(t *testing.T) {
	inputText := "text"
	testCases := []struct {
		marker Marker
		output string
	}{
		{
			marker: Marker{
				Color: color.Color{
					Background: "red",
					Charactor:  "blue",
				},
				Format: format.Format{
					Bold:          true,
					Italic:        true,
					Underline:     true,
					Hide:          true,
					Strikethrough: true,
				},
			},
			output: "\x1b[41m\x1b[34m\x1b[1m\x1b[3m\x1b[4m\x1b[8m\x1b[9m" + inputText + "\x1b[0m",
		},
		{
			marker: Marker{
				Color: color.Color{
					Background: "",
					Charactor:  "",
				},
				Format: format.Format{
					Bold:          false,
					Italic:        false,
					Underline:     false,
					Hide:          false,
					Strikethrough: false,
				},
			},
			output: inputText,
		},
	}

	for _, tc := range testCases {
		expected := tc.output
		actual := tc.marker.MarkText(inputText)
		if actual != expected {
			msg := "Did not add correctory style code to head and tail of text.\n"
			t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
		}
	}
}
