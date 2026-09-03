package figure

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func NewCircle(r float64) (Circle, error) {
	if r <= 0 {
		return Circle{}, fmt.Errorf("wrong params: radius = %f", r)
	}

	return Circle{radius: r}, nil
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}
