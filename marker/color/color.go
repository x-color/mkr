package color

import (
	"fmt"
	"strconv"
)

// colorToNumber maps color name to number
var colorToNumber = map[string]int{
	"black":   0,
	"blue":    4,
	"cyan":    6,
	"green":   2,
	"magenta": 5,
	"red":     1,
	"yellow":  3,
}

type Color struct {
	Background string
	Charactor  string
}

// GenColorCode generates background and charactor color code to add to head of
// text.
func (color Color) GenColorCode() string {
    colorCode := genBackColor(color.Background)
    colorCode += genCharColor(color.Charactor)
    return colorCode
}

// genBackColor generates background color code from color name or number.
func genBackColor(colorString string) string {
    normalColor := genNormalBackColorCode(colorString)
    extendedColor := genExtendedBackColorCode(colorString)
    return normalColor + extendedColor
}

// genNormalBackColorCode returns 8 colors of background.
// It is suppurted by most terminals.
func genNormalBackColorCode(colorName string) string {
	if num, ok := colorToNumber[colorName]; ok {
		return fmt.Sprintf("\x1b[4%dm", num)
	}
	return ""
}

// genExtendedBackColorCode returns string of 256 colors of background.
// It is suppurted by some terminals.
func genExtendedBackColorCode(colorNumber string) string {
	if num, err := strconv.Atoi(colorNumber); err == nil {
		if 0 <= num && num <= 255 {
			return fmt.Sprintf("\x1b[48;5;%dm", num)
		}
	}
	return ""
}

// genCharColor generates charactor color code from color name or number.
func genCharColor(colorString string) string {
    normalColor := genNormalCharColorCode(colorString)
    extendedColor := genExtendedCharColorCode(colorString)
    return normalColor + extendedColor
}

// genNormalCharColorCode returns 8 colors of charactor.
// It is suppurted by most terminals.
func genNormalCharColorCode(colorName string) string {
	if num, ok := colorToNumber[colorName]; ok {
		return fmt.Sprintf("\x1b[3%dm", num)
	}
	return ""
}

// genExtendedCharColorCode returns string of 256 colors of charactor.
// It is suppurted by some terminals.
func genExtendedCharColorCode(colorNumber string) string {
	if num, err := strconv.Atoi(colorNumber); err == nil {
		if 0 <= num && num <= 255 {
			return fmt.Sprintf("\x1b[38;5;%dm", num)
		}
	}
	return ""
}
