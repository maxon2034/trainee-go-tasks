package figure

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(w, h float64) (Rectangle, error) {
	if w <= 0 || h <= 0 {
		return Rectangle{}, fmt.Errorf("wrong params: width = %f, height = %f", w, h)
	}

	return Rectangle{width: w, height: h}, nil
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
