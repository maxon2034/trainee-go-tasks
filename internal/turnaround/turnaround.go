package turnaround

import "strings"

func Do(s string) string {
	var word []rune
	var turnaround []rune
	res := make([]rune, 0, len([]rune(s)))

	for i, v := range s {
		if s[i] == ' ' {
			continue
		}

		word = append(word, v)

		if i+1 == len(s) || s[i+1] == ' ' {
			for i := range word {
				turnaround = append(turnaround, word[len(word)-1-i])
			}
			res = append(res, turnaround...)
			res = append(res, ' ')

			word = word[:0]
			turnaround = turnaround[:0]
		}
	}
	return string(res)
}

func DoLib(s string) string {
	split := strings.Split(s, " ")
	var turnaround []rune
	res := make([]string, 0, len(split))
	for _, v := range split {
		for l := range v {
			turnaround = append(turnaround, []rune(v)[len([]rune(v))-1-l])
		}
		res = append(res, string(turnaround))
		turnaround = turnaround[:0]
	}

	return strings.Join(res, " ")
}
