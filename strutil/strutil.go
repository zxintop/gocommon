package strutil

import (
	"strings"
	"unicode"
)

// ToUpper 将字符串进行大写
func ToUpper(s string) string {
	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToUpper(r[i])
	}
	return string(r)
}

func ToLower(s string) string {
	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToLower(r[i])
	}
	return string(r)
}

func Split(s string, sep string) []string {
	return strings.Split(s, sep)
}
