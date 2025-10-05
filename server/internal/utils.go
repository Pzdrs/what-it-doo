package internal

import "unicode"

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s) // convert string to runes
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
