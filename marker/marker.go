package marker

import (
	"github.com/x-color/mkr/marker/color"
	"github.com/x-color/mkr/marker/format"
)

// Marker has marking style, color and format.
type Marker struct {
	Color  color.Color
	Format format.Format
}

// MarkText adds style code to head and tail of text.
func (marker Marker) MarkText(text string) string {
	styleCode := marker.Color.GenColorCode()
	styleCode += marker.Format.GenFormatCode()
	if styleCode == "" {
		return text
	}
	return styleCode + text + "\x1b[0m"
}
