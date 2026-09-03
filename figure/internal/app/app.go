package app

import (
	"fmt"
	"log"

	"github.com/maxon2034/trainee-go-tasks/figure/internal/figure"
)

func Run(shape string, r, w, h float64) {
	var s figure.Shape
	var err error
	switch shape {
	case "circle":
		s, err = figure.NewCircle(r)
		if err != nil {
			log.Fatal(err)
		}
	case "rectangle":
		s, err = figure.NewRectangle(w, h)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Area: %.2f", s.Area())
}
