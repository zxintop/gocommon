package strutil

import (
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
