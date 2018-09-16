package reader

import (
	"bufio"
	"io"
)

var Input io.Reader

// ReadFile reads text in files and send it to channel.
func ReadFile(file io.Reader, lines chan string) {
	s := bufio.NewScanner(file)
	for s.Scan() {
		lines <- s.Text()
	}
	close(lines)
}
