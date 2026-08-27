package app

import (
	"fmt"

	"github.com/maxon2034/trainee-go-tasks/palindrome/internal/palindrome"
)

func Run() {
	s := []string{
		"Mr. Owl ate my metal worm",
		"22/2/22",
		"Довод",
		"Повод",
	}
	for _, v := range s {
		if palindrome.Is(v) {
			fmt.Println(v, "is a palindrome")
		} else {
			fmt.Println(v, "is not a palindrome")
		}
	}

}
