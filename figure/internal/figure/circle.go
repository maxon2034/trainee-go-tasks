package figure

import "math"

type Circle struct {
	radius float64
}

func (c *Circle) Set(r float64) {
	c.radius = r
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}
