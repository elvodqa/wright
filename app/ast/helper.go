package ast

import "unicode"

func isLetter(b byte) bool {
	r := rune(b)
	return unicode.IsLetter(r)
}

func isDigit(b byte) bool {
	r := rune(b)
	return unicode.IsDigit(r)
}
