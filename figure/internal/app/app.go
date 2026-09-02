package app

import (
	"fmt"

	"github.com/maxon2034/trainee-go-tasks/figure/internal/figure"
)

func Run() {
	c := figure.Rectangle{}
	res, err := figure.CalcArea(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
