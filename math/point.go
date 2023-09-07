package math

type Point [2]float64

func NewPoint(vec Vector2) Point {
	return Point{vec.X(), vec.Y()}
}

func (p *Point) X() float64 {
	return p[0]
}
func (p *Point) Y() float64 {
	return p[1]
}

func (p *Point) Vector2() Vector2 {
	return Vector2{p.X(), p.Y()}
}
