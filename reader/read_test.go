package reader

import (
	"bytes"
	"testing"
)

func TestReadFiles(t *testing.T) {
	file := bytes.NewBufferString("hello\nworld\ntesting now")
	output := []string{"hello", "world", "testing now"}
	lines := make(chan string, 100)

	go ReadFile(file, lines)

	i := 0
	for actual := range lines {
		expected := output[i]
		if actual != expected {
			msg := "Read and outputed text is not expected text."
			t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, expected, actual)
		}
		i++
	}

	if i != len(output) {
		msg := "Did not read all text in files.\n"
		t.Fatalf("%s\nExpected: %#v\nActual  : %#v\n", msg, i, len(output))
	}
}
