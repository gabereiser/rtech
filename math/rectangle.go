package math

type Rectangle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewRectangle(x, y, width, height float64) *Rectangle {
	return &Rectangle{
		x, y, width, height,
	}
}
