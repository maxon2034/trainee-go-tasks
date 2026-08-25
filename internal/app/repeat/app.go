package app

import (
	"fmt"

	"github.com/maxon2034/trainee-go-tasks/internal/repeat"
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
	res := repeat.Remove(s)

	fmt.Println(" after:", res)
}
