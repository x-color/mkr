package color

import "testing"

func TestGenBackColor(t *testing.T) {
	testCases := []struct {
		color  Color
		output string
	}{
		{
			color: Color{
				Background: "red",
				Charactor:  "blue",
			},
			output: "\x1b[41m\x1b[34m",
		},
		{
			color: Color{
				Background: "",
				Charactor:  "",
			},
			output: "",
		},
		{
			color: Color{
				Background: "1",
				Charactor:  "255",
			},
			output: "\x1b[48;5;1m\x1b[38;5;255m",
		},
	}
	for _, tc := range testCases {
		expected := tc.output
		actual := tc.color.GenColorCode()
		if actual != expected {
			msg := "Did not generate correct color code.\n"
			t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
		}
	}
}
