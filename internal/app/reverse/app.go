package app

import (
	"fmt"

	"github.com/maxon2034/trainee-go-tasks/internal/reverse"
)

func Run() {
	s := "уауау ваики"
	fmt.Println("before: ", s)
	fmt.Println("no lib: ", reverse.Words(s))
	fmt.Println("with lib: ", reverse.WordsLib(s))
}
