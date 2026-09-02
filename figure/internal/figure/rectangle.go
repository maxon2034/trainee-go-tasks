package figure

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Set(w float64, h float64) {
	r.width = w
	r.height = h
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
