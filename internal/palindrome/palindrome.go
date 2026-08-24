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
	for i, v := range []rune(s) {
		if len([]rune(s))-i >= i && v == []rune(s)[(len([]rune(s))-1)-i] {
			if (len([]rune(s))-1)-i == i {
				break
			}
		} else {
			return false
		}
	}
	return true
}
