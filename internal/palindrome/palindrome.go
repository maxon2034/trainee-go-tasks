package palindrome

import (
	"strings"
	"unicode"
)

func Is(s string) bool {
	s = strings.ToLower(s)
	runes := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			runes = append(runes, r)
		}
	}
	left, right := 0, len(runes)-1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}

	return true
}
