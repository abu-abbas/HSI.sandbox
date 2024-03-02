package str

import "strings"

func LeftPad(text string, length int, pad string) string {
	textLen := len([]rune(text))

	if length <= textLen {
		return text
	}

	var pattern string
	if pad == "" {
		pattern = " "
	} else {
		pattern = pad
	}

	repeat := length - textLen
	result := strings.Repeat(pattern, repeat) + text

	return result
}
