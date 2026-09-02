package figure

import (
	"flag"
	"fmt"
)

type Shape interface {
	Area() float64
}

func CalcArea(s Shape) (float64, error) {

	shape := flag.String("shape", "", "shape")
	width := flag.Float64("width", 0, "width")
	height := flag.Float64("height", 0, "height")
	radius := flag.Float64("radius", 0, "radius")

	flag.Parse()

	switch *shape {
	case "circle":
		c := s.(Circle)

		if *radius < 0 {
			return 0, fmt.Errorf("Error: wrong radius parameter")
		}

		c.Set(*radius)
		return c.Area(), nil
	case "rectangle":
		r := s.(Rectangle)

		if *width < 0 {
			return 0, fmt.Errorf("Error: wrong width parameter")
		}

		if *height < 0 {
			return 0, fmt.Errorf("Error: wrong height parameter")
		}

		r.Set(*width, *height)
		return r.Area(), nil
	default:
		return 0, fmt.Errorf("unknown shape")
	}
}
