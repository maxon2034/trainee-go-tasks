package app

import (
	"fmt"

	"github.com/maxon2034/trainee-go-tasks/internal/duplicate"
)

func Run() {
	s := []string{
		"aaa",
		"bbb",
		"ccc",
		"aaa",
		"ddd",
		"ddd",
	}
	fmt.Println("before:", s)
	res := duplicate.Remove(s)

	fmt.Println(" after:", res)
}
