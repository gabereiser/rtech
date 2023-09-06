package math

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x, y,
	}
}
