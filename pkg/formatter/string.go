package formatter

import "unicode"

func LowerCaseFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func UpperCaseFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}
