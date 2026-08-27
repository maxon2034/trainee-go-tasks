package app

import (
	"fmt"

	"github.com/maxon2034/trainee-go-tasks/duplicate/internal/duplicate"
)

func Run() {
	ss := []string{
		"aaa",
		"bbb",
		"ccc",
		"aaa",
		"ddd",
		"ddd",
	}
	fmt.Println("before:", ss)
	res := duplicate.Remove(ss)

	fmt.Println(" after:", res)
}
