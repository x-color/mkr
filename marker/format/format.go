package format

// Format has some text format
type Format struct {
	Bold          bool
	Italic        bool
	Underline     bool
	Hide          bool
	Strikethrough bool
}

// GenFormatCode generates text format code
func (format Format) GenFormatCode() string {
	formatCode := genBoldFormat(format.Bold)
	formatCode += genItalicFormat(format.Italic)
	formatCode += genUnderlineFormat(format.Underline)
	formatCode += genHideFormat(format.Hide)
	formatCode += genStrikethroughFormat(format.Strikethrough)
	return formatCode
}

// genBoldFormat generates bold format code
func genBoldFormat(yes bool) string {
	if yes {
		return "\x1b[1m"
	}
	return ""
}

// genItalicFormat generates italic format code
func genItalicFormat(yes bool) string {
	if yes {
		return "\x1b[3m"
	}
	return ""
}

// genUnderlineFormat generates underline format code
func genUnderlineFormat(yes bool) string {
	if yes {
		return "\x1b[4m"
	}
	return ""
}

// genHideFormat generates hide format code
func genHideFormat(yes bool) string {
	if yes {
		return "\x1b[8m"
	}
	return ""
}

// genStrikethroughFormat generates strikethrough format code
func genStrikethroughFormat(yes bool) string {
	if yes {
		return "\x1b[9m"
	}
	return ""
}
