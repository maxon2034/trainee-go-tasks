package reverse

import "strings"

func Words(s string) string {
	var word []rune
	var reverse []rune

	runes := []rune(s)
	res := make([]rune, 0, len(runes))

	for i, v := range runes {
		if runes[i] == ' ' {
			continue
		}

		word = append(word, v)

		if i+1 == len(runes) || runes[i+1] == ' ' {
			for i := range word {
				reverse = append(reverse, word[len(word)-1-i])
			}
			res = append(res, reverse...)
			res = append(res, ' ')

			word = word[:0]
			reverse = reverse[:0]
		}
	}
	return string(res)
}

func WordsLib(s string) string {
	ss := strings.Split(s, " ")
	var runes []rune
	var reverse []rune

	res := make([]string, 0, len(ss))
	for _, v := range ss {
		runes = []rune(v)
		for l := range runes {
			reverse = append(reverse, runes[len(runes)-1-l])
		}
		res = append(res, string(reverse))
		reverse = reverse[:0]
	}

	return strings.Join(res, " ")
}
