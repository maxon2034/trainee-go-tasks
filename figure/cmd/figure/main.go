package main

import (
	"flag"

	"github.com/maxon2034/trainee-go-tasks/figure/internal/app"
)

func main() {

	shape := flag.String("shape", "", "shape")
	width := flag.Float64("width", 0, "width")
	height := flag.Float64("height", 0, "height")
	radius := flag.Float64("radius", 0, "radius")

	flag.Parse()

	app.Run(*shape, *radius, *width, *height)
}
