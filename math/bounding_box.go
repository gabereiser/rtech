package math

type BoundingBox struct {
	Origin Vector3
	Min    Vector3
	Max    Vector3
}

type BoundingBoxAxisAligned struct {
	Origin      Vector3
	Min         Vector3
	Max         Vector3
	Orientation Quaternion
}
