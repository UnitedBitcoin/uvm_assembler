package uvmassembler

import (
	"strings"
)

func Trim(str string) string {
	return strings.Trim(str, " \t\r\n")
}

func StringFirstIndexOf(str string, fn func(byte) bool) int {
	for i := 0; i < len(str); i++ {
		c := str[i]
		if fn(c) {
			return i
		}
	}
	return len(str)
}

func StringFirstIndexOfNot(str string, fn func(byte) bool) int {
	for i := 0; i < len(str); i++ {
		c := str[i]
		if !fn(c) {
			return i
		}
	}
	return len(str)
}

func IsEmptyChar(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n'
}
