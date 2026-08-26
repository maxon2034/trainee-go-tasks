package app

import (
	"fmt"

	"github.com/maxon2034/trainee-go-tasks/internal/turnaround"
)

func Run() {
	s := "abs adsf sdv wefwef seff"
	fmt.Println("before: ", s)
	fmt.Println("no lib: ", turnaround.Do(s))
	fmt.Println("with lib: ", turnaround.DoLib(s))
}
