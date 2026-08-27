package palindrome

import (
	"strings"
)

func Is(s string) bool {
	split := strings.Split(s, " ")
	for i, v := range split {
		split[i] = strings.Trim(v, ".,?!")
	}
	s = strings.Join(split, "")
	s = strings.ToLower(s)

	runes := []rune(s)

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
