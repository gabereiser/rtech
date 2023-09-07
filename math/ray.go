package math

type Ray struct {
	Origin    *Vector3
	Direction *Vector3
}

func NewRay(origin *Vector3, direction *Vector3) *Ray {
	return &Ray{origin, direction}
}
