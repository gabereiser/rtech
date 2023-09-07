package math

type Rectangle [4]float64

func NewRectangle(x, y, width, height float64) Rectangle {
	return Rectangle{
		x, y, width, height,
	}
}
func (r *Rectangle) X() float64 {
	return r[0]
}
func (r *Rectangle) Y() float64 {
	return r[1]
}
func (r *Rectangle) Width() float64 {
	return r[2]
}
func (r *Rectangle) Height() float64 {
	return r[3]
}

func (r *Rectangle) Left() float64 {
	return r.X()
}

func (r *Rectangle) Right() float64 {
	return r.X() + r.Width()
}

func (r *Rectangle) Top() float64 {
	return r.Y()
}

func (r *Rectangle) Bottom() float64 {
	return r.Y() + r.Height()
}

func (r *Rectangle) IsEmpty() bool {
	return (r.X() == 0 && r.Y() == 0 && r.Width() == 0 && r.Height() == 0)
}

func (r *Rectangle) Position() Point {
	return NewPoint(Vector2{r.X(), r.Y()})
}

func (r *Rectangle) Center() Point {
	return NewPoint(Vector2{r.X() + (r.Width() / 2), r.Y() + (r.Height() / 2)})
}

func (r *Rectangle) PointWithin(p *Point) bool {
	return (p.X() > r.X() && p.X() < (r.X()+r.Width())) && (p.Y() > r.Y() && p.Y() < (r.Y()+r.Height()))
}

func (r *Rectangle) Intersects(rect *Rectangle) bool {
	return rect.Left() < r.Right() &&
		r.Left() < rect.Right() &&
		rect.Top() < r.Bottom() &&
		r.Top() < rect.Bottom()
}
